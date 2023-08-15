package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceDetailByIdRequest Request Object
type ShowWorkspaceDetailByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 依据workspace id查工作区
	ModelId string `json:"model_id"`
}

func (o ShowWorkspaceDetailByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceDetailByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceDetailByIdRequest", string(data)}, " ")
}
