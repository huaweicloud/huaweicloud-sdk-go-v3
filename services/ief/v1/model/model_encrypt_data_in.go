package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 加密数据信息
type EncryptDataIn struct {

	// 加密数据名称
	Name string `json:"name"`

	// 加密数据描述
	Description string `json:"description"`

	// 加密数据项配置
	Config []EncryptDataItem `json:"config"`
}

func (o EncryptDataIn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataIn struct{}"
	}

	return strings.Join([]string{"EncryptDataIn", string(data)}, " ")
}
