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
type ShowPublicipTagsResponse struct {
	// 标签列表
	Tags []ResourceTagResp `json:"tags,omitempty"`
}

func (o ShowPublicipTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPublicipTagsResponse", string(data)}, " ")
}
