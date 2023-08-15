package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasourceConnectionsRequest Request Object
type ListDatasourceConnectionsRequest struct {

	// 标签
	Tags *string `json:"tags,omitempty"`
}

func (o ListDatasourceConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListDatasourceConnectionsRequest", string(data)}, " ")
}
