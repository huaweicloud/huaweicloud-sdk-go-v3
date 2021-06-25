package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLogContents3Request struct {
	// 日志流id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsStructLogParamsNew `json:"body,omitempty"`
}

func (o UpdateLogContents3Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContents3Request struct{}"
	}

	return strings.Join([]string{"UpdateLogContents3Request", string(data)}, " ")
}
