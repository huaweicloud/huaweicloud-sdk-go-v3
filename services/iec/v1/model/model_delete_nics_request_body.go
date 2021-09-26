package model

import (
	"encoding/json"

	"strings"
)

// 删除网卡请求体。
type DeleteNicsRequestBody struct {
	// 待删除网卡信息。

	Nics []NicId `json:"nics"`
}

func (o DeleteNicsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNicsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteNicsRequestBody", string(data)}, " ")
}