package public

import (
	"path/filepath"
	"runtime"
)

func Path() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	return basePath
}
