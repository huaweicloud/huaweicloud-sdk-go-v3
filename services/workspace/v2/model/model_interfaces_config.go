package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterfacesConfig 接口配置。
type InterfacesConfig struct {

	// 应用平台。
	ApplyPlatform *string `json:"apply_platform,omitempty"`

	// ID。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 结果。
	Results map[string]string `json:"results,omitempty"`

	// 分页信息。
	Pagination map[string]interface{} `json:"pagination,omitempty"`

	Request *InterfacesRequest `json:"request,omitempty"`

	// 响应。
	Response *string `json:"response,omitempty"`

	// 检查结果。
	ResultCheck *string `json:"result_check,omitempty"`
}

func (o InterfacesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfacesConfig struct{}"
	}

	return strings.Join([]string{"InterfacesConfig", string(data)}, " ")
}
