package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataTransformationRequest Request Object
type CreateDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateDataTransformationRequestBody `json:"body,omitempty"`
}

func (o CreateDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"CreateDataTransformationRequest", string(data)}, " ")
}
