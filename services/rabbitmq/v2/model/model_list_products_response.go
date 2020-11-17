/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProductsResponse struct {
	// 表示按需付费的产品列表。
	Hourly *[]ListProductsRespHourly `json:"Hourly,omitempty"`
	// 表示包年包月的产品列表。当前暂不支持通过API创建包年包月的rabbitmq实例。
	Monthly *[]ListProductsRespHourly `json:"Monthly,omitempty"`
}

func (o ListProductsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
