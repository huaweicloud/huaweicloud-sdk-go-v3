package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunDeletePictureResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDeletePictureResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunDeletePictureResponse struct{}"
	}

	return strings.Join([]string{"RunDeletePictureResponse", string(data)}, " ")
}
