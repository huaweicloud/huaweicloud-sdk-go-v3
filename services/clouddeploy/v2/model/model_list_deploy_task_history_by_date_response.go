package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDeployTaskHistoryByDateResponse struct {

	// 部署任务历史执行记录列表
	Result *[]ExecuteRecordV2Body `json:"result,omitempty" xml:"result"`

	// 开始时间和结束时间内任务历史执行记录总数
	TotalNum       *int32 `json:"total_num,omitempty" xml:"total_num"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDeployTaskHistoryByDateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployTaskHistoryByDateResponse struct{}"
	}

	return strings.Join([]string{"ListDeployTaskHistoryByDateResponse", string(data)}, " ")
}
