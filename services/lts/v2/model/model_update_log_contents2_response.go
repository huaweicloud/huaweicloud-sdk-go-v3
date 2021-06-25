package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLogContents2Response struct {
	// 日志信息。

	StructLogs     *[]StructLogContents `json:"struct_logs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateLogContents2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContents2Response struct{}"
	}

	return strings.Join([]string{"UpdateLogContents2Response", string(data)}, " ")
}
