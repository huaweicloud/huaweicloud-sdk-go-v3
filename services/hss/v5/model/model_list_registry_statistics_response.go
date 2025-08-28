package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryStatisticsResponse Response Object
type ListRegistryStatisticsResponse struct {

	// **参数解释**： 接入异常数量 **取值范围**： 0-100
	FailNum *int32 `json:"fail_num,omitempty"`

	// **参数解释**： 接入成功数量 **取值范围**： 0-100
	SuccessNum     *int32 `json:"success_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRegistryStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegistryStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListRegistryStatisticsResponse", string(data)}, " ")
}
