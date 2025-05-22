package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDatabaseOmUserActionResponse Response Object
type ExecuteDatabaseOmUserActionResponse struct {

	// **参数解释**： 错误码。非0皆为异常场景。 **取值范围**： 不涉及。
	ErrorCode *int32 `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	OmUserInfo     *DatabaseOmUserInfo `json:"om_user_info,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ExecuteDatabaseOmUserActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDatabaseOmUserActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDatabaseOmUserActionResponse", string(data)}, " ")
}
