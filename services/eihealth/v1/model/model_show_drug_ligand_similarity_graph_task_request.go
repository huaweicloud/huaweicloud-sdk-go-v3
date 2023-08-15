package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrugLigandSimilarityGraphTaskRequest Request Object
type ShowDrugLigandSimilarityGraphTaskRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 任务ID，通过创建任务接口取得。
	TaskId string `json:"task_id"`
}

func (o ShowDrugLigandSimilarityGraphTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrugLigandSimilarityGraphTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowDrugLigandSimilarityGraphTaskRequest", string(data)}, " ")
}
