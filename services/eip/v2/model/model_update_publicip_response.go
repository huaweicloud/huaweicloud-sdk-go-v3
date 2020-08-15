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
type UpdatePublicipResponse struct {
	Publicip *PublicipShowResp `json:"publicip,omitempty"`
}

func (o UpdatePublicipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipResponse", string(data)}, " ")
}
