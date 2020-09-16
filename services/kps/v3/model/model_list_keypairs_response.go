/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListKeypairsResponse struct {
	// SSH密钥对信息详情
	Keypairs *[]interface{} `json:"keypairs,omitempty"`
}

func (o ListKeypairsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeypairsResponse", string(data)}, " ")
}
