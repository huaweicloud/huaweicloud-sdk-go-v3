package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunCheckTaskJobsResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *[]CheckTaskJobsItemsBody `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RunCheckTaskJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCheckTaskJobsResponse struct{}"
	}

	return strings.Join([]string{"RunCheckTaskJobsResponse", string(data)}, " ")
}
