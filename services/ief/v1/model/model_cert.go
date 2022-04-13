package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 证书属性配置
type Cert struct {
	// 证书名称

	Name string `json:"name"`
	// 证书描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Description *string `json:"description,omitempty"`
	// 证书类型，包含： - application：应用证书 - device：设备证书

	Type string `json:"type"`
}

func (o Cert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cert struct{}"
	}

	return strings.Join([]string{"Cert", string(data)}, " ")
}
