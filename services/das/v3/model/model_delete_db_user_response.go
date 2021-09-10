package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDbUserResponse struct {
	// 删除结果，删除成功返回OK

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDbUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbUserResponse", string(data)}, " ")
}
