package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeployTaskHistoryByDateResponse Response Object
type ListDeployTaskHistoryByDateResponse struct {

	// 应用历史部署记录列表
	Result *[]ExecuteRecordV2Body `json:"result,omitempty"`

	// 开始时间和结束时间内应用历史部署记录总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDeployTaskHistoryByDateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployTaskHistoryByDateResponse struct{}"
	}

	return strings.Join([]string{"ListDeployTaskHistoryByDateResponse", string(data)}, " ")
}
