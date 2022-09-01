package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListReadWriteRatioRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分页参数：起始值 [大于等于0] 。
	CurPage string `json:"curPage" xml:"curPage"`

	// 分页参数：每页多少条。
	PerPage string `json:"perPage" xml:"perPage"`

	// 开始时间，UTC time，精确到毫秒。
	StartDate string `json:"startDate" xml:"startDate"`

	// 结束时间，UTC time，精确到毫秒。结束时间与开始时间，间隔不能超过1个月。
	EndDate string `json:"endDate" xml:"endDate"`
}

func (o ListReadWriteRatioRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadWriteRatioRequest struct{}"
	}

	return strings.Join([]string{"ListReadWriteRatioRequest", string(data)}, " ")
}
