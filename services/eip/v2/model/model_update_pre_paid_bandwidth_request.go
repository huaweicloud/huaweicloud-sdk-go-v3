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
type UpdatePrePaidBandwidthRequest struct {
	BandwidthId string                             `json:"bandwidth_id"`
	Body        *UpdatePrePaidBandwidthRequestBody `json:"body,omitempty"`
}

func (o UpdatePrePaidBandwidthRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePrePaidBandwidthRequest", string(data)}, " ")
}
