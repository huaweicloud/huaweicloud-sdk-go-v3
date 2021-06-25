package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLogContents2Request struct {
	// 日志组id。

	LogGroupId string `json:"log_group_id"`
	// 日志流id。

	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsStructLogParams `json:"body,omitempty"`
}

func (o UpdateLogContents2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContents2Request struct{}"
	}

	return strings.Join([]string{"UpdateLogContents2Request", string(data)}, " ")
}
