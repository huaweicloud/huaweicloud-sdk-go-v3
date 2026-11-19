package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindServerRouteResponse Response Object
type BindServerRouteResponse struct {

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindServerRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindServerRouteResponse struct{}"
	}

	return strings.Join([]string{"BindServerRouteResponse", string(data)}, " ")
}
