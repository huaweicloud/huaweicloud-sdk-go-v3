package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLogContentsRequest struct {
	// 日志组id。

	LogGroupId string `json:"log_group_id"`
	// 日志流id。

	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsLogParams `json:"body,omitempty"`
}

func (o UpdateLogContentsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContentsRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogContentsRequest", string(data)}, " ")
}
