package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PathMap path:文件路径
type PathMap struct {

	// 文件路径
	Path string `json:"path"`
}

func (o PathMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PathMap struct{}"
	}

	return strings.Join([]string{"PathMap", string(data)}, " ")
}
