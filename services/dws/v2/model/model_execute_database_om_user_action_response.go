package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDatabaseOmUserActionResponse Response Object
type ExecuteDatabaseOmUserActionResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
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
