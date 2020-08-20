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

// Request Object
type ListPublicipsByTagsRequest struct {
	Body *ListPublicipsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListPublicipsByTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublicipsByTagsRequest", string(data)}, " ")
}
