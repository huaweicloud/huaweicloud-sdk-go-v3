package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCpiJobRequest Request Object
type CreateCpiJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateCpiJobReq `json:"body,omitempty"`
}

func (o CreateCpiJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCpiJobRequest struct{}"
	}

	return strings.Join([]string{"CreateCpiJobRequest", string(data)}, " ")
}
