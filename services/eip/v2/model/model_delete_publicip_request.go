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
type DeletePublicipRequest struct {
	PublicipId string `json:"publicip_id"`
}

func (o DeletePublicipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePublicipRequest", string(data)}, " ")
}
