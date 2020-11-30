/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 更新裸金属服务器元数据
type MetaData struct {
	Metadata *KeyValue `json:"metadata"`
}

func (o MetaData) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MetaData", string(data)}, " ")
}
