package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBcsEventsResponse Response Object
type ListBcsEventsResponse struct {

	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListBcsEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsEventsResponse struct{}"
	}

	return strings.Join([]string{"ListBcsEventsResponse", string(data)}, " ")
}
