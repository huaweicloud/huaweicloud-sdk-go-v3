package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBcsEventsStatisticResponse Response Object
type ListBcsEventsStatisticResponse struct {

	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListBcsEventsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsEventsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListBcsEventsStatisticResponse", string(data)}, " ")
}
