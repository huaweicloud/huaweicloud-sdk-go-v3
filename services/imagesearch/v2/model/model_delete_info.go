package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInfo 删除结果的相关信息。
type DeleteInfo struct {

	// 符合条件的结果总数。
	TotalNum float32 `json:"total_num,omitempty"`

	// 本次删除的结果总数，目前一次请求最多删除100条结果。
	DeleteNum float32 `json:"delete_num,omitempty"`
}

func (o DeleteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInfo struct{}"
	}

	return strings.Join([]string{"DeleteInfo", string(data)}, " ")
}
