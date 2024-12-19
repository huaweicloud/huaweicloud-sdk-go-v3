package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalDcGatewayRequest Request Object
type ShowGlobalDcGatewayRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`

	// 根据企业项目ID过滤资源实例
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`
}

func (o ShowGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalDcGatewayRequest", string(data)}, " ")
}
