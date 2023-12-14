package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueriesRequest Request Object
type ListQueriesRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *ListQueriesRequestBody `json:"body,omitempty"`
}

func (o ListQueriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesRequest struct{}"
	}

	return strings.Join([]string{"ListQueriesRequest", string(data)}, " ")
}
