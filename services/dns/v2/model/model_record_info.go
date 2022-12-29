package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordInfo struct {

	// 找回记录host名称。
	Host *string `json:"host,omitempty"`

	// 找回记录解析值。
	Value *string `json:"value,omitempty"`
}

func (o RecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfo struct{}"
	}

	return strings.Join([]string{"RecordInfo", string(data)}, " ")
}
