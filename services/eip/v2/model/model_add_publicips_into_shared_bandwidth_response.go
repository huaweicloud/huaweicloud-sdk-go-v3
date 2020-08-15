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
type AddPublicipsIntoSharedBandwidthResponse struct {
	Bandwidth *BandwidthRespInsert `json:"bandwidth,omitempty"`
}

func (o AddPublicipsIntoSharedBandwidthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddPublicipsIntoSharedBandwidthResponse", string(data)}, " ")
}
