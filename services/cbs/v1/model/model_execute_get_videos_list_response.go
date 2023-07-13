package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetVideosListResponse Response Object
type ExecuteGetVideosListResponse struct {

	// 配额
	Quota int32 `json:"quota"`

	// 总数
	Total int32 `json:"total"`

	// 偏移
	Offset int32 `json:"offset"`

	// 返回数量
	Count int32 `json:"count"`

	// 视频列表
	Videos         []Video `json:"videos"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteGetVideosListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetVideosListResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGetVideosListResponse", string(data)}, " ")
}
