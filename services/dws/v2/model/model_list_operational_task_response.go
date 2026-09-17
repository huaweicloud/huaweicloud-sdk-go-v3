package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationalTaskResponse Response Object
type ListOperationalTaskResponse struct {

	// **参数解释**： 总条数。 **默认取值**： 大于等于0。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 任务列表。 **默认取值**： 大于等于0。
	Data           *[]TaskInfoVo `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListOperationalTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationalTaskResponse struct{}"
	}

	return strings.Join([]string{"ListOperationalTaskResponse", string(data)}, " ")
}
