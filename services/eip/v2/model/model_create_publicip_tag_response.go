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
type CreatePublicipTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePublicipTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePublicipTagResponse", string(data)}, " ")
}
