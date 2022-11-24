package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type Show3dStructureContentRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// study_id
	StudyId string `json:"study_id"`

	// study作业id
	JobId string `json:"job_id"`

	// 配体名称
	Ligand string `json:"ligand"`

	// 受体名称
	Receptor string `json:"receptor"`
}

func (o Show3dStructureContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Show3dStructureContentRequest struct{}"
	}

	return strings.Join([]string{"Show3dStructureContentRequest", string(data)}, " ")
}
