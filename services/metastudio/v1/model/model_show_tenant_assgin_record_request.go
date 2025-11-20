package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantAssginRecordRequest Request Object
type ShowTenantAssginRecordRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowTenantAssginRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantAssginRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantAssginRecordRequest", string(data)}, " ")
}
