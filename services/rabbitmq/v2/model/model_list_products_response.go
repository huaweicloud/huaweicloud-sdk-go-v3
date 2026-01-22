package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsResponse Response Object
type ListProductsResponse struct {

	// **参数解释**： 表示按需付费的产品列表。
	Hourly *[]ListProductsRespHourly `json:"Hourly,omitempty"`

	// **参数解释**： 表示包年包月的产品列表。当前暂不支持通过API创建包年包月的Kafka实例。
	Monthly        *[]ListProductsRespMonthly `json:"Monthly,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
