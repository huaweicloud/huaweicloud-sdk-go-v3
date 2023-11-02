package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPocketMolDesignJobRequest Request Object
type ShowPocketMolDesignJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowPocketMolDesignJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPocketMolDesignJobRequest struct{}"
	}

	return strings.Join([]string{"ShowPocketMolDesignJobRequest", string(data)}, " ")
}
