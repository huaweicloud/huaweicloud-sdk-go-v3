package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PagedInfo 分页返回信息
type PagedInfo struct {

	// 本次返回的对象个数
	CurrentCount int32 `json:"current_count"`

	// 下一页查询地址
	NextMarker string `json:"next_marker"`

	// 上一页查询地址
	PreviousMarker string `json:"previous_marker"`
}

func (o PagedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PagedInfo struct{}"
	}

	return strings.Join([]string{"PagedInfo", string(data)}, " ")
}
