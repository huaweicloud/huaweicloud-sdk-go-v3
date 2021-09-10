package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLinkResponse struct {
	// 校验结构：如果创建连接失败，返回失败原因，请参见validation-result参数说明。如果创建成功，返回空列表。

	ValidationResult *[]ValidationResult `json:"validation-result,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o UpdateLinkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLinkResponse struct{}"
	}

	return strings.Join([]string{"UpdateLinkResponse", string(data)}, " ")
}
