package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrugModelRequest Request Object
type ListDrugModelRequest struct {

	// 模糊搜索值
	SearchKey *string `json:"search_key,omitempty"`

	// 创建者列表
	CreatorList *[]string `json:"creator_list,omitempty"`

	// 模型类型列表
	TypeList *[]string `json:"type_list,omitempty"`

	// 模型状态列表
	StatusList *[]string `json:"status_list,omitempty"`

	// 排序规则 目前默认时间降序，支持根据create_time|finish_time|base_model_name
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则 目前默认时间降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 最小创建时间
	CreateStartTime *int64 `json:"create_start_time,omitempty"`

	// 最大创建时间
	CreateEndTime *int64 `json:"create_end_time,omitempty"`

	// 最小结束时间
	FinishStartTime *int64 `json:"finish_start_time,omitempty"`

	// 最大结束时间
	FinishEndTime *int64 `json:"finish_end_time,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 基模型id列表
	BaseModelList *[]string `json:"base_model_list,omitempty"`
}

func (o ListDrugModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugModelRequest struct{}"
	}

	return strings.Join([]string{"ListDrugModelRequest", string(data)}, " ")
}
