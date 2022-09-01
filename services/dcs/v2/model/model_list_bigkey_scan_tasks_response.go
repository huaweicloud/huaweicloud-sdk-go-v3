package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBigkeyScanTasksResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 大key分析记录列表
	Records        *[]RecordsResponse `json:"records,omitempty" xml:"records"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBigkeyScanTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigkeyScanTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBigkeyScanTasksResponse", string(data)}, " ")
}
