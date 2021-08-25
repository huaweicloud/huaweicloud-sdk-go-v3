package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunAddPictureResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunAddPictureResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunAddPictureResponse struct{}"
	}

	return strings.Join([]string{"RunAddPictureResponse", string(data)}, " ")
}
