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
type ListKeypairDetailResponse struct {
	Keypair *KeypairDetail `json:"keypair,omitempty"`
}

func (o ListKeypairDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeypairDetailResponse", string(data)}, " ")
}
