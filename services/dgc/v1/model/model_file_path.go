package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilePath struct {

	// 文件在OBS上的路径
	Path string `json:"path"`
}

func (o FilePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilePath struct{}"
	}

	return strings.Join([]string{"FilePath", string(data)}, " ")
}
