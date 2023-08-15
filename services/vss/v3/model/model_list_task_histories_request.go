package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskHistoriesRequest Request Object
type ListTaskHistoriesRequest struct {

	// 网站域名ID
	DomainId string `json:"domain_id"`

	// 分页查询，偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTaskHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTaskHistoriesRequest", string(data)}, " ")
}
