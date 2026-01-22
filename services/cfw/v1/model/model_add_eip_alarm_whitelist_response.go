package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEipAlarmWhitelistResponse Response Object
type AddEipAlarmWhitelistResponse struct {

	// **参数解释**： 添加EIP告警白名单响应data **取值范围**： 不涉及
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o AddEipAlarmWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipAlarmWhitelistResponse struct{}"
	}

	return strings.Join([]string{"AddEipAlarmWhitelistResponse", string(data)}, " ")
}
