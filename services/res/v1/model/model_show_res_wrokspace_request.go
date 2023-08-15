package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResWrokspaceRequest Request Object
type ShowResWrokspaceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowResWrokspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResWrokspaceRequest struct{}"
	}

	return strings.Join([]string{"ShowResWrokspaceRequest", string(data)}, " ")
}
