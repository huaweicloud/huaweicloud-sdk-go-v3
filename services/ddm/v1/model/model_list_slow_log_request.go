package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSlowLogRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 分页参数：起始值 [大于等于0] 。

	CurPage string `json:"curPage"`
	// 分页参数：每页多少条。

	PerPage string `json:"perPage"`
	// 开始时间，UTC time，精确到毫秒。

	StartDate string `json:"startDate"`
	// 结束时间，UTC time，精确到毫秒。

	EndDate string `json:"endDate"`
}

func (o ListSlowLogRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogRequest", string(data)}, " ")
}
