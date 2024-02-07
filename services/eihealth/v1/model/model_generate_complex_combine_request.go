package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateComplexCombineRequest Request Object
type GenerateComplexCombineRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *GenerateComplexCombineReq `json:"body,omitempty"`
}

func (o GenerateComplexCombineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateComplexCombineRequest struct{}"
	}

	return strings.Join([]string{"GenerateComplexCombineRequest", string(data)}, " ")
}
