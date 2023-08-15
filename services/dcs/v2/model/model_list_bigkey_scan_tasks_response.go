package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBigkeyScanTasksResponse Response Object
type ListBigkeyScanTasksResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 大key分析记录列表
	Records        *[]RecordsResponse `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBigkeyScanTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigkeyScanTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBigkeyScanTasksResponse", string(data)}, " ")
}
