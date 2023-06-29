package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBindApiRequest Request Object
type SearchBindApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// app编号
	AppId string `json:"app_id"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o SearchBindApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBindApiRequest struct{}"
	}

	return strings.Join([]string{"SearchBindApiRequest", string(data)}, " ")
}
