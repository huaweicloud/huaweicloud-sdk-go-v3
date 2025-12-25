package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectConfigRequest Request Object
type ListCollectConfigRequest struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序
	SortDir *string `json:"sort_dir,omitempty"`

	// 云服务
	Csvc *string `json:"csvc,omitempty"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 是否查询云服务接入指标， 默认为 True
	QueryStatistics *bool `json:"query_statistics,omitempty"`
}

func (o ListCollectConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigRequest struct{}"
	}

	return strings.Join([]string{"ListCollectConfigRequest", string(data)}, " ")
}
