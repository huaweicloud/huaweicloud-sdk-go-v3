package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelsRequest Request Object
type ListModelsRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	// 每页数量，默认10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分组ID筛选。
	GroupId *string `json:"group_id,omitempty"`

	// 名称模糊搜索。
	Name *string `json:"name,omitempty"`
}

func (o ListModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelsRequest struct{}"
	}

	return strings.Join([]string{"ListModelsRequest", string(data)}, " ")
}
