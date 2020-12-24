/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ServiceResourceInfo struct {
	BasicInfo *ResourceBasicInfo `json:"basic_info,omitempty"`
}

func (o ServiceResourceInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ServiceResourceInfo", string(data)}, " ")
}
