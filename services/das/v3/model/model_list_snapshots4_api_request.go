package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshots4ApiRequest Request Object
type ListSnapshots4ApiRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 模块
	Module *int32 `json:"module,omitempty"`

	// 开始时间（Unix时间戳，毫秒）
	StartAt *int64 `json:"start_at,omitempty"`

	// 结束时间（Unix时间戳，毫秒）
	EndAt *int64 `json:"end_at,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`
}

func (o ListSnapshots4ApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshots4ApiRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshots4ApiRequest", string(data)}, " ")
}
