package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsInt64(a []int64, x int64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func GetRequestId (ctx context.Context) string {
	var requestID string
	request := ctx.Value("requestID")
	if request != nil {
		requestID = request.(string)
	} else {
		requestID = "no_context"
	}
	return requestID
}


func RandomString(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Println(err)
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	rs = rs[0 : len(rs)-1]

	return rs
}
