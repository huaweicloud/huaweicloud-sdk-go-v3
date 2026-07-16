package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHyperinstanceTagsRequest Request Object
type CreateHyperinstanceTagsRequest struct {

	// **参数解释**：Lite Server 超节点ID。
	Id string `json:"id"`

	Body *TagRequest `json:"body,omitempty"`
}

func (o CreateHyperinstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHyperinstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateHyperinstanceTagsRequest", string(data)}, " ")
}
