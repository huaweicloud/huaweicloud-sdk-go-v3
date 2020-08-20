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
type DeletePublicipResponse struct {
}

func (o DeletePublicipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePublicipResponse", string(data)}, " ")
}
