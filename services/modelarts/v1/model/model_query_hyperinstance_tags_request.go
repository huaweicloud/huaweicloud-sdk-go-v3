package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryHyperinstanceTagsRequest Request Object
type QueryHyperinstanceTagsRequest struct {

	// **参数解释**：Lite Server 超节点ID。
	Id string `json:"id"`
}

func (o QueryHyperinstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryHyperinstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"QueryHyperinstanceTagsRequest", string(data)}, " ")
}
