package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

const FileName = "mock.json"

type Mock struct {
	Id          string          `json:"id"`
	Url         string          `json:"mainUrl"`
	Method      string          `json:"method"`
	Params      []*RequestParam `json:"params"`
	Status      string          `json:"status"`
	ContentType string          `json:"contentType"`
	Headers		[]*Header		`json:"headers"`
	Body        string          `json:"body"`
}

type Mocks []*Mock

type RequestParam struct {
	Name   string `json:"name"`
	Regexp string `json:"regexp"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetConfigMap() (map[string][]*Mock, error) {
	c, err := readJsonConfig()
	if err != nil {
		return nil, err
	}
	cm := ConfigMap(c)

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

func RemoveFromConfig(id string) (Mocks, bool, error) {
	isDeleted := false
	mocks, err := readJsonConfig()
	if err != nil {
		return nil, false, err
	}
	for i, m := range mocks {
		if m.Id == id {
			mocks = append(mocks[:i], mocks[i+1:]...)
			isDeleted = true
		}
	}

	err = writeConfig(mocks)
	if err != nil {
		return nil, false, err
	}

	return mocks, isDeleted, nil
}

func UpdateConfig(mock *Mock) (Mocks, bool, error) {
	isUpdated := false
	mocks, err := readJsonConfig()
	if err != nil {
		return nil, false, err
	}
	for i, m := range mocks {
		if m.Id == mock.Id {
			mocks[i] = mock
			isUpdated = true
		}
	}

	err = writeConfig(mocks)
	if err != nil {
		return nil, false, err
	}

	return mocks, isUpdated, nil
}

func GetConfigFileBody() ([]byte, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	defer func() {
		configFile.Close()
	}()
	byteValue, err := ioutil.ReadAll(configFile)

	return byteValue, err
}

func writeConfig(mocks Mocks) error {

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
	byteValue, err := GetConfigFileBody()
	if err != nil {
		return nil, err
	}

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

func ConfigMap(mocks Mocks) map[string][]*Mock {
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
