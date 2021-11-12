package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfo struct {
	// 返回下一页查询地址(为空时，代表后面没有数据)

	NextMarker string `json:"next_marker"`
	// 返回前一页查询地址

	PreviousMarker string `json:"previous_marker"`
	// 本页返回条目数量

	CurrentCount string `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
