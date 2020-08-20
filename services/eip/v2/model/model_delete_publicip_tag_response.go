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
type DeletePublicipTagResponse struct {
}

func (o DeletePublicipTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePublicipTagResponse", string(data)}, " ")
}
