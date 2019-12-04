package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

const FileName = "mock.json"

type Mock struct {
	Url         string          `json:"url"`
	Method      string          `json:"method"`
	Params      []*RequestParam `json:"params"`
	Status      int             `json:"status"`
	ContentType string          `json:"content_type"`
	Body        string          `json:"body"`
}

type Mocks []*Mock

type RequestParam struct {
	Name   string `json:"name"`
	Regexp string `json:"regexp"`
}

func GetConfigMap() (map[string][]*Mock, error) {
	c, err := readJsonConfig()
	if err != nil {
		return nil, err
	}
	cm := configMap(c)

	return cm, nil
}

func AddToConfig(m *Mock) error {
	mocks, err := readJsonConfig()
	if err != nil {
		return err
	}

	mocks = append(mocks, m)

	err = writeConfig(mocks)

	return err
}

func RemoveFromConfig (url, method string) (bool, error) {
	isDeleted := false

	mocks, err := readJsonConfig()
	if err != nil {
		return false, err
	}

	for i, m := range mocks {
		if m.Method == method && strings.Trim(m.Url, "/") == strings.Trim(url, "/") {
			mocks = append(mocks[:i], mocks[i+1:]...)
			isDeleted = true
		}
	}

	err = writeConfig(mocks)
	if err != nil {
		return false, err
	}

	return isDeleted, nil
}

func writeConfig (mocks Mocks) error {

	configFile, err := getConfigFile()
	if err != nil {
		return err
	}

	defer func() {
		configFile.Close()
	}()

	jsonBytes, err := json.Marshal(mocks)
	if err != nil {
		return err
	}

	err = configFile.Truncate(0)
	if err != nil {
		return err
	}
	_, err = configFile.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func readJsonConfig() (Mocks, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	defer func() {
		configFile.Close()
	}()

	byteValue, _ := ioutil.ReadAll(configFile)

	mocks := &Mocks{}
	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, mocks)
		if err != nil {
			return nil, err
		}
	}

	return *mocks, nil
}

func getConfigFile() (*os.File, error) {
	_, err := os.Stat(FileName)
	if os.IsNotExist(err) {
		_, err = os.Create(FileName)
		if err != nil {
			return nil, err
		}
	}

	configFile, err := os.OpenFile(FileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return configFile, nil
}

func configMap(mocks Mocks) map[string][]*Mock {
	configMap := map[string][]*Mock{}
	for _, m := range mocks {
		url := strings.Trim(m.Url, "/")
		if len(url) < 1 {
			continue
		}
		sm, ok := configMap[url]
		if ok {
			configMap[url] = append(sm, m)
		} else {
			configMap[url] = []*Mock{m}
		}
	}

	return configMap
}
