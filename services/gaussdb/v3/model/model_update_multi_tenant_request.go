package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMultiTenantRequest Request Object
type UpdateMultiTenantRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *OperateMultiTenantReq `json:"body,omitempty"`
}

func (o UpdateMultiTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiTenantRequest struct{}"
	}

	return strings.Join([]string{"UpdateMultiTenantRequest", string(data)}, " ")
}
