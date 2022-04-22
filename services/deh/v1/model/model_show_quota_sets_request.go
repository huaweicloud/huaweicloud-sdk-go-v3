package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowQuotaSetsRequest struct {

	// 租户ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。
	TenantId string `json:"tenant_id"`

	// 配额类别。
	Resource *string `json:"resource,omitempty"`
}

func (o ShowQuotaSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaSetsRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaSetsRequest", string(data)}, " ")
}
