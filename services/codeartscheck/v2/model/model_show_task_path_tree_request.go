package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskPathTreeRequest Request Object
type ShowTaskPathTreeRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 任务id
	TaskId string `json:"task_id"`

	// 目录或文件的路径
	CurrentPath *string `json:"current_path,omitempty"`

	// 分页索引，偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量,每页最多显示1000条
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTaskPathTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskPathTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskPathTreeRequest", string(data)}, " ")
}
