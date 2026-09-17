package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationalTaskDetailResponse Response Object
type ListOperationalTaskDetailResponse struct {

	// **参数解释**： 详细列表。 **默认取值**： 0
	Data *[]TaskStatusOpenResp `json:"data,omitempty"`

	// **参数解释**： 总条数。 **默认取值**： 0
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOperationalTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationalTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ListOperationalTaskDetailResponse", string(data)}, " ")
}
