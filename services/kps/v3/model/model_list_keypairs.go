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

// SSH密钥对信息详情
type ListKeypairs struct {
	Keypair *Keypair `json:"keypair,omitempty"`
}

func (o ListKeypairs) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeypairs", string(data)}, " ")
}
