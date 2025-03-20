package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Description 描述信息。
type Description struct {

	// 英文描述。
	EnUS string `json:"en_US"`

	// 中文描述。
	ZhCN string `json:"zh_CN"`
}

func (o Description) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Description struct{}"
	}

	return strings.Join([]string{"Description", string(data)}, " ")
}
