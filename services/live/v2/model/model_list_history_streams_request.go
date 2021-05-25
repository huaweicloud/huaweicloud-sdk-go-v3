package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListHistoryStreamsRequest struct {
	// 推流域名。

	Domain string `json:"domain"`
	// 应用名称。

	App *string `json:"app,omitempty"`
	// 分页编号，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数。  取值范围：[1,100]  默认值：10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHistoryStreamsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHistoryStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryStreamsRequest", string(data)}, " ")
}
