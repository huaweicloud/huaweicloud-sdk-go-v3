package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 分页参数：起始值 [大于等于0] 。

	Offset int32 `json:"offset"`
	// 分页参数：每页多少条 [大于0且小于等于128]。

	Limit int32 `json:"limit"`
}

func (o ListInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
