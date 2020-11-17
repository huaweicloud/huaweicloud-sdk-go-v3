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
type ListPluginsResponse struct {
	// 插件信息列表。
	Plugins *[]ListPluginsRespPlugins `json:"plugins,omitempty"`
}

func (o ListPluginsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPluginsResponse", string(data)}, " ")
}
