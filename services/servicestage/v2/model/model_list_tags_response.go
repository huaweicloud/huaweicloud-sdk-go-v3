/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTagsResponse struct {
	// 项目tag标签列表。
	Tags *[]string `json:"tags,omitempty"`
}

func (o ListTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
