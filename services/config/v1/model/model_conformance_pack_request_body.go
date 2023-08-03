package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackRequestBody 创建合规规则包的请求体。
type ConformancePackRequestBody struct {

	// 合规规则包名称。
	Name string `json:"name"`

	// 委托名称，该委托需要授权云服务ResourceFormation调用Config服务的合规规则的创建、更新、删除接口。
	AgencyName *string `json:"agency_name,omitempty"`

	// 预定义合规包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 自定义合规包内容。
	TemplateBody *string `json:"template_body,omitempty"`

	// 合规包模板OBS地址。
	TemplateUri *string `json:"template_uri,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`
}

func (o ConformancePackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackRequestBody struct{}"
	}

	return strings.Join([]string{"ConformancePackRequestBody", string(data)}, " ")
}
