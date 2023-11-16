package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectAsyncStatusLogInfoResponse Response Object
type ShowProjectAsyncStatusLogInfoResponse struct {

	// 日志组名称
	GroupName *string `json:"group_name,omitempty"`

	// 日志组id
	GroupId *string `json:"group_id,omitempty"`

	// 日志流id
	StreamId *string `json:"stream_id,omitempty"`

	// 日志流名称
	StreamName     *string `json:"stream_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectAsyncStatusLogInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectAsyncStatusLogInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectAsyncStatusLogInfoResponse", string(data)}, " ")
}
