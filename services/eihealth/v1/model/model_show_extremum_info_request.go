package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtremumInfoRequest Request Object
type ShowExtremumInfoRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// study_id
	StudyId string `json:"study_id"`

	// study作业id
	JobId string `json:"job_id"`
}

func (o ShowExtremumInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtremumInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowExtremumInfoRequest", string(data)}, " ")
}
