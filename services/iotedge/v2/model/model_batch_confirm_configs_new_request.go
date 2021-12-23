package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchConfirmConfigsNewRequest struct {
	// 边缘节点ID

	NodeId string `json:"node_id"`
	// 边侧第三方应用的模块ID

	IaId string `json:"ia_id"`

	Body *ConfirmIaConfigsRequestBody `json:"body,omitempty"`
}

func (o BatchConfirmConfigsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmConfigsNewRequest struct{}"
	}

	return strings.Join([]string{"BatchConfirmConfigsNewRequest", string(data)}, " ")
}
