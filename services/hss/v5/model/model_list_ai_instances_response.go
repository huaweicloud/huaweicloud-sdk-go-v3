package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiInstancesResponse Response Object
type ListAiInstancesResponse struct {

	// **参数解释**： AI智能体实例的总数 **取值范围**： 0~2000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： AI智能体实例列表 **取值范围**： 不涉及
	DataList       *[]ListAiInstancesResponseBodyData `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListAiInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAiInstancesResponse", string(data)}, " ")
}
