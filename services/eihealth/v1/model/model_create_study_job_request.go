package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStudyJobRequest Request Object
type CreateStudyJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// study_id
	StudyId string `json:"study_id"`

	Body *CreateStudyJobReq `json:"body,omitempty"`
}

func (o CreateStudyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStudyJobRequest struct{}"
	}

	return strings.Join([]string{"CreateStudyJobRequest", string(data)}, " ")
}
