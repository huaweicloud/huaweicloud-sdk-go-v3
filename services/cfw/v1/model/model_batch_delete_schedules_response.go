package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSchedulesResponse Response Object
type BatchDeleteSchedulesResponse struct {

	// **参数解释**： 被删除的时间表ID列表 **取值范围**： 不涉及
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteSchedulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSchedulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSchedulesResponse", string(data)}, " ")
}
