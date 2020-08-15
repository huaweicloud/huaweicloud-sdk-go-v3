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
type BatchDeletePublicipTagsResponse struct {
}

func (o BatchDeletePublicipTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeletePublicipTagsResponse", string(data)}, " ")
}
