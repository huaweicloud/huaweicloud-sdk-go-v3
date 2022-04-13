package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchImportConfigsRequest struct {
	// 边缘节点ID

	NodeId string `json:"node_id"`
	// 边侧第三方应用的模块ID

	IaId string `json:"ia_id"`

	Body *BatchImportConfigsRequestBody `json:"body,omitempty"`
}

func (o BatchImportConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchImportConfigsRequest", string(data)}, " ")
}
