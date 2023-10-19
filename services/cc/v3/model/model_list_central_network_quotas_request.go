package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkQuotasRequest Request Object
type ListCentralNetworkQuotasRequest struct {

	// 根据配额类型查询，可查询多个类型。
	QuotaType *[]CentralNetworkQuotaKeyEnum `json:"quota_type,omitempty"`

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`
}

func (o ListCentralNetworkQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkQuotasRequest", string(data)}, " ")
}
