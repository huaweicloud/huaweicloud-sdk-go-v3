package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMysqlInstanceRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *MysqlInstanceRequest `json:"body,omitempty"`
}

func (o CreateMysqlInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateMysqlInstanceRequest", string(data)}, " ")
}
