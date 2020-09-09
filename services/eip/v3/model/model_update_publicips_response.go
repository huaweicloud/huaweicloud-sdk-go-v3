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
type UpdatePublicipsResponse struct {
	// 本次请求的编号
	RequestId *string           `json:"request_id,omitempty"`
	Publicip  *PublicipShowResp `json:"publicip,omitempty"`
}

func (o UpdatePublicipsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipsResponse", string(data)}, " ")
}
