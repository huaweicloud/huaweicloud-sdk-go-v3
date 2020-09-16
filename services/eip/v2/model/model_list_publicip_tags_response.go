/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPublicipTagsResponse struct {
	// 标签列表
	Tags *[]TagResp `json:"tags,omitempty"`
}

func (o ListPublicipTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublicipTagsResponse", string(data)}, " ")
}
