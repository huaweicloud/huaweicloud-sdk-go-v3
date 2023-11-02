package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePocketMolDesignJobRequest Request Object
type CreatePocketMolDesignJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreatePocketMolDesignJobReq `json:"body,omitempty"`
}

func (o CreatePocketMolDesignJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketMolDesignJobRequest struct{}"
	}

	return strings.Join([]string{"CreatePocketMolDesignJobRequest", string(data)}, " ")
}
