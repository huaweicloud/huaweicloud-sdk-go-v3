package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobResp 创建任务响应体
type CreateJobResp struct {

	// 任务ID
	Id string `json:"id"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务状态 CREATING：创建中，CREATE_FAILED：创建失败，CONFIGURATION：配置中，WAITING_FOR_START：等待启动中，RELEASE_RESOURCE_COMPLETE：已结束，DELETED：已删除，INCRE_TRANSFER_STARTED：增量迁移中，INCRE_TRANSFER_FAILED：增量迁移失败，FULL_TRANSFER_STARTED：全量迁移中，FULL_TRANSFER_COMPLETE：全量迁移完成，PAUSING：暂停中，FULL_TRANSFER_FAILED：全量迁移失败
	Status *string `json:"status,omitempty"`

	// 创建时间，时间戳
	CreateTime *string `json:"create_time,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 子任务ID，有子任务时返回该字段。
	ChildIds *[]string `json:"child_ids,omitempty"`
}

func (o CreateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResp struct{}"
	}

	return strings.Join([]string{"CreateJobResp", string(data)}, " ")
}
