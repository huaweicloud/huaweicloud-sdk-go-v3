package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyReqBody The request body of update agency request.
type UpdateAgencyReqBody struct {

	// 委托或信任委托最大会话时长，默认为3600秒。
	MaxSessionDuration *int32 `json:"max_session_duration,omitempty"`

	// 委托或信任委托描述信息。
	Description *string `json:"description,omitempty"`
}

func (o UpdateAgencyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyReqBody struct{}"
	}

	return strings.Join([]string{"UpdateAgencyReqBody", string(data)}, " ")
}
