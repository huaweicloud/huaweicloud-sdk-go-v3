package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowUserInstancesResponse struct {
	// 非默认用户主密钥个数。

	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowUserInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesResponse", string(data)}, " ")
}
