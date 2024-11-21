package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFormatConverterRequest Request Object
type RunFormatConverterRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RunFormatConverterReq `json:"body,omitempty"`
}

func (o RunFormatConverterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFormatConverterRequest struct{}"
	}

	return strings.Join([]string{"RunFormatConverterRequest", string(data)}, " ")
}
