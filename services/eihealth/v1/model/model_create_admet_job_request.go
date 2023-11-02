package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAdmetJobRequest Request Object
type CreateAdmetJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateAdmetJobReq `json:"body,omitempty"`
}

func (o CreateAdmetJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdmetJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAdmetJobRequest", string(data)}, " ")
}
