package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRepairDetailResp 修复详情。
type QueryRepairDetailResp struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 修复详情。
	RepairDetails *[]QueryRepairDetailRespRepairDetails `json:"repair_details,omitempty"`
}

func (o QueryRepairDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRepairDetailResp struct{}"
	}

	return strings.Join([]string{"QueryRepairDetailResp", string(data)}, " ")
}
