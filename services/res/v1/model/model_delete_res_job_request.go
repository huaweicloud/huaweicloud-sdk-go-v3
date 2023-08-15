package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResJobRequest Request Object
type DeleteResJobRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 资源id（数据源id或场景id）
	ResourceId string `json:"resource_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o DeleteResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteResJobRequest", string(data)}, " ")
}
