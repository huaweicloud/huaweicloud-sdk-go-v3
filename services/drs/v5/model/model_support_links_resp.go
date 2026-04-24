package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportLinksResp 查询可用链路信息响应体
type SupportLinksResp struct {

	// 列表中的项目总数，与分页无关。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 支持的链路
	SupportLinks *[]SupportLinkInfo `json:"support_links,omitempty"`
}

func (o SupportLinksResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportLinksResp struct{}"
	}

	return strings.Join([]string{"SupportLinksResp", string(data)}, " ")
}
