package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricRelationsResponse Response Object
type ListMetricRelationsResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListMetricRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricRelationsResponse", string(data)}, " ")
}
