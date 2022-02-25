package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建后端服务器响应结果
type BatchDeleteMemberState struct {
	// 后端服务器ID。

	Id string `json:"id"`
	// 当前后端服务器删除结果状态。取值： - successful：删除成功。 - not found：member不存在。

	RetStatus string `json:"ret_status"`
}

func (o BatchDeleteMemberState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMemberState struct{}"
	}

	return strings.Join([]string{"BatchDeleteMemberState", string(data)}, " ")
}
