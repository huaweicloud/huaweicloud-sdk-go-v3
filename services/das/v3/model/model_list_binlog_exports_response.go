package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogExportsResponse Response Object
type ListBinlogExportsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 导出任务列表
	TaskList       *[]BinlogExportTaskDetailResp `json:"task_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListBinlogExportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogExportsResponse struct{}"
	}

	return strings.Join([]string{"ListBinlogExportsResponse", string(data)}, " ")
}
