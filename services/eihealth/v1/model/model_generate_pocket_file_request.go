package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneratePocketFileRequest Request Object
type GeneratePocketFileRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RunPocketReq `json:"body,omitempty"`
}

func (o GeneratePocketFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneratePocketFileRequest struct{}"
	}

	return strings.Join([]string{"GeneratePocketFileRequest", string(data)}, " ")
}
