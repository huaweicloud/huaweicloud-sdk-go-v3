package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelVersionRequest Request Object
type DeleteModelVersionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// Service ID
	ModelId string `json:"model_id"`

	// Version ID
	VersionId string `json:"version_id"`
}

func (o DeleteModelVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteModelVersionRequest", string(data)}, " ")
}
