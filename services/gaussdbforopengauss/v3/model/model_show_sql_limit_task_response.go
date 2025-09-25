package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitTaskResponse Response Object
type ShowSqlLimitTaskResponse struct {

	// **参数解释**: 限流任务拦截次数。 **取值范围**: 不涉及。
	LimitCount     *int32 `json:"limit_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitTaskResponse", string(data)}, " ")
}
