package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceTypes struct {

	// 云服务类型的名称。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 云服务类型的编码。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务类型的缩写。
	Abbreviation *string `json:"abbreviation,omitempty"`
}

func (o ServiceTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceTypes struct{}"
	}

	return strings.Join([]string{"ServiceTypes", string(data)}, " ")
}
