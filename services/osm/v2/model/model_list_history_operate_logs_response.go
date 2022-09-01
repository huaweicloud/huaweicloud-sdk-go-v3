package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryOperateLogsResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 操作列表
	OpsList        *[]OperateLog `json:"ops_list,omitempty" xml:"ops_list"`
	HttpStatusCode int           `json:"-"`
}

func (o ListHistoryOperateLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOperateLogsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryOperateLogsResponse", string(data)}, " ")
}
