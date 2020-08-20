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
type CreatePublicipResponse struct {
	Publicip *PublicipCreateResp `json:"publicip,omitempty"`
}

func (o CreatePublicipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePublicipResponse", string(data)}, " ")
}
