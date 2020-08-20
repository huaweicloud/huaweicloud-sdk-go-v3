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
type RemovePublicipsFromSharedBandwidthRequest struct {
	BandwidthId string                                         `json:"bandwidth_id"`
	Body        *RemovePublicipsFromSharedBandwidthRequestBody `json:"body,omitempty"`
}

func (o RemovePublicipsFromSharedBandwidthRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemovePublicipsFromSharedBandwidthRequest", string(data)}, " ")
}
