package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DifficultInfo 习题难度信息
type DifficultInfo struct {

	// 难度id
	Id string `json:"id"`

	// 难度名称
	Name string `json:"name"`

	// 难度等级
	Degree int32 `json:"degree"`
}

func (o DifficultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DifficultInfo struct{}"
	}

	return strings.Join([]string{"DifficultInfo", string(data)}, " ")
}
