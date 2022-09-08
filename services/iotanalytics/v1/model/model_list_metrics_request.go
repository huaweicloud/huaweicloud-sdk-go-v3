package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMetricsRequest struct {

	// 存储ID
	DataStoreId string `json:"data_store_id"`

	Body *GetMetricsRequest `json:"body,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}
