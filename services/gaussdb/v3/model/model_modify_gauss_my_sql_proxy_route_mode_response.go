package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyGaussMySqlProxyRouteModeResponse Response Object
type ModifyGaussMySqlProxyRouteModeResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyGaussMySqlProxyRouteModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGaussMySqlProxyRouteModeResponse struct{}"
	}

	return strings.Join([]string{"ModifyGaussMySqlProxyRouteModeResponse", string(data)}, " ")
}
