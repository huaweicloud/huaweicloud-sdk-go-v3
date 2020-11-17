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
type ShowInstanceExtendProductInfoResponse struct {
	// 表示按需付费的产品列表。
	Hourly *[]ListProductsRespHourly `json:"hourly,omitempty"`
	// 表示包年包月的产品列表。当前暂不支持通过API创建包年包月的Rabbitmq实例。
	Monthly *[]ListProductsRespHourly `json:"monthly,omitempty"`
}

func (o ShowInstanceExtendProductInfoResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceExtendProductInfoResponse", string(data)}, " ")
}
