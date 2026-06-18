package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantCmksRequest Request Object
type ListTenantCmksRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTenantCmksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantCmksRequest struct{}"
	}

	return strings.Join([]string{"ListTenantCmksRequest", string(data)}, " ")
}
