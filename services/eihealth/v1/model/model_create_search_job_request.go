package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchJobRequest Request Object
type CreateSearchJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateSearchJobReq `json:"body,omitempty"`
}

func (o CreateSearchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSearchJobRequest", string(data)}, " ")
}
