package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobRequest Request Object
type BatchDeleteJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *BatchOperateJobReq `json:"body,omitempty"`
}

func (o BatchDeleteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobRequest", string(data)}, " ")
}
