package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageScanTaskResponse Response Object
type ListImageScanTaskResponse struct {

	// **参数解释**: task任务列表 **取值范围**: 最小值0，最大值1000
	DataList *[]ImageScanTaskInfo `json:"data_list,omitempty"`

	// **参数解释**: 任务总数 **取值范围**: 最小值0，最大值1000
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImageScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageScanTaskResponse struct{}"
	}

	return strings.Join([]string{"ListImageScanTaskResponse", string(data)}, " ")
}
