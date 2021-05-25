package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVersionRequest struct {
	// 待查询版本号。当前仅支持v2。

	ApiVersion string `json:"api_version"`
}

func (o ListVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVersionRequest struct{}"
	}

	return strings.Join([]string{"ListVersionRequest", string(data)}, " ")
}
