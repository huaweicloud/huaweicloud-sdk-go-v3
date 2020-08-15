/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowUserInstancesResponse struct {
	// 非默认用户主密钥个数。
	InstanceNum string `json:"instance_num,omitempty"`
}

func (o ShowUserInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUserInstancesResponse", string(data)}, " ")
}
