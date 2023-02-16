package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 检查结果的相关信息。
type CheckInfo struct {

	// 符合条件的结果总数。
	TotalNum float32 `json:"total_num,omitempty"`

	// 返回的结果总数。
	ReturnNum float32 `json:"return_num,omitempty"`

	LastItem *SearchAfterParam `json:"last_item,omitempty"`
}

func (o CheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckInfo struct{}"
	}

	return strings.Join([]string{"CheckInfo", string(data)}, " ")
}
