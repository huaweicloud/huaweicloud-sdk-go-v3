package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterTimezoneResponse Response Object
type ModifyClusterTimezoneResponse struct {

	// **参数解释**： 错误码。 **取值范围**： 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**： 任务ID，可用于查询任务进度信息。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyClusterTimezoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterTimezoneResponse struct{}"
	}

	return strings.Join([]string{"ModifyClusterTimezoneResponse", string(data)}, " ")
}
