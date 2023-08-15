package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAuthorizeAppRequest Request Object
type SearchAuthorizeAppRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// api编号
	ApiId string `json:"api_id"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o SearchAuthorizeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAuthorizeAppRequest struct{}"
	}

	return strings.Join([]string{"SearchAuthorizeAppRequest", string(data)}, " ")
}
