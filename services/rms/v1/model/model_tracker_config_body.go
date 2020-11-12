/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// TrackerConfig对象
type TrackerConfigBody struct {
	Channel  *ChannelConfigBody  `json:"channel"`
	Selector *SelectorConfigBody `json:"selector"`
	// IAM委托名称
	AgencyName string `json:"agency_name"`
}

func (o TrackerConfigBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TrackerConfigBody", string(data)}, " ")
}
