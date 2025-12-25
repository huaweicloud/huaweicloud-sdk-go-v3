package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIndicatorsRequest Request Object
type ExportIndicatorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ExportIndicatorsRequestBody `json:"body,omitempty"`
}

func (o ExportIndicatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIndicatorsRequest struct{}"
	}

	return strings.Join([]string{"ExportIndicatorsRequest", string(data)}, " ")
}
