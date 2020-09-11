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

// 解绑密钥对描述消息体
type DisassociateKeypairRequestBody struct {
	Server *DisassociateEcsServerInfo `json:"server"`
}

func (o DisassociateKeypairRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisassociateKeypairRequestBody", string(data)}, " ")
}
