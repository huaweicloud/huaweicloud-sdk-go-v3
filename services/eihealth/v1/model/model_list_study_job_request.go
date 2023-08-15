package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStudyJobRequest Request Object
type ListStudyJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// study_id
	StudyId string `json:"study_id"`
}

func (o ListStudyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStudyJobRequest struct{}"
	}

	return strings.Join([]string{"ListStudyJobRequest", string(data)}, " ")
}
