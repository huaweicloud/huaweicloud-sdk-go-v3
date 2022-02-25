package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	// 结束时间，UTC time，精确到毫秒。结束时间与开始时间，间隔不能超过7天。

	EndDate string `json:"endDate"`
}

func (o ListSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogRequest", string(data)}, " ")
}
