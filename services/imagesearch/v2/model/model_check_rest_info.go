package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 检查数据的相关信息。
type CheckRestInfo struct {

	// 数据是否存在，存在返回true，不存在返回false。仅在指定ID检查时包含该字段。
	Existed *bool `json:"existed,omitempty"`

	ItemInfo *ItemSource `json:"item_info,omitempty"`

	// 检查结果列表，仅在条件检查时包含该字段。
	Items *[]SearchItem `json:"items,omitempty"`

	CheckInfo *CheckInfo `json:"check_info,omitempty"`
}

func (o CheckRestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRestInfo struct{}"
	}

	return strings.Join([]string{"CheckRestInfo", string(data)}, " ")
}
