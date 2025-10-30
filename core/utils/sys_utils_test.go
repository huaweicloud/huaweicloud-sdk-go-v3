package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestGetOSInfo(t *testing.T) {
//	info := getOSInfo()
//	assert.NotEmpty(t, info)
//}

func TestGetEnvInfoString(t *testing.T) {
	cloudDirName = t.TempDir()
	info := GetEnvInfoString()
	assert.NotEmpty(t, info)
}
