/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ListStreamForbiddenResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 禁播黑名单列表
	Blocks []StreamForbiddenList `json:"blocks,omitempty"`
}

func (o ListStreamForbiddenResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStreamForbiddenResponse", string(data)}, " ")
}
