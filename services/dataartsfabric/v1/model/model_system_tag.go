package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemTag 系统标签。
type SystemTag struct {

	// 键，固定值。
	Key string `json:"key"`

	// 企业项目ID。
	Value string `json:"value"`
}

func (o SystemTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemTag struct{}"
	}

	return strings.Join([]string{"SystemTag", string(data)}, " ")
}
