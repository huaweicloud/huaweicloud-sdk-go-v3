package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageReq 分页请求体
type PageReq struct {

	// 当前页, 不能超过item除每页任务数量的最大页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页item数量，填0获取全部item
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o PageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageReq struct{}"
	}

	return strings.Join([]string{"PageReq", string(data)}, " ")
}
