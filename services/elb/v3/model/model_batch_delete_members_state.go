package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建后端服务器响应结果
type BatchDeleteMembersState struct {

	// 后端服务器ID。 >说明： 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id" xml:"id"`

	// 当前后端服务器删除结果状态。取值： - successful：删除成功。 - not found：member不存在。
	RetStatus string `json:"ret_status" xml:"ret_status"`
}

func (o BatchDeleteMembersState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersState struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersState", string(data)}, " ")
}
