package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRequest Request Object
type CreateAlertRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAlertRequestBody `json:"body,omitempty"`
}

func (o CreateAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertRequest", string(data)}, " ")
}
