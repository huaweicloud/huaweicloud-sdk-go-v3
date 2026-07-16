package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperinstanceTagsRequest Request Object
type DeleteHyperinstanceTagsRequest struct {

	// **参数解释**：Lite Server 超节点ID。
	Id string `json:"id"`

	Body *TagRequest `json:"body,omitempty"`
}

func (o DeleteHyperinstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperinstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteHyperinstanceTagsRequest", string(data)}, " ")
}
