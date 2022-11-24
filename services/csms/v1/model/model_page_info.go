package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页信息。
type PageInfo struct {

	// 下一页查询地址（本页的末尾凭据名称，下一页起始凭据名称）。
	NextMarker *string `json:"next_marker,omitempty"`

	// 本页的起始凭据名称，上一页末尾凭据名称。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 本页返回条目数量。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
