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
type ShowPublicipTagsRequest struct {
	PublicipId string `json:"publicip_id"`
}

func (o ShowPublicipTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPublicipTagsRequest", string(data)}, " ")
}
