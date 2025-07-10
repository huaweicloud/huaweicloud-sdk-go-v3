package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppUserAccessData struct {

	// 用户名称。
	Username *string `json:"username,omitempty"`

	// 接入失败数。
	AccessFailedCount *int32 `json:"access_failed_count,omitempty"`

	// 最近一次接入失败时间，UTC时间，格式为：2022-05-11T11:45:42Z。
	LastAccessFailedTime *string `json:"last_access_failed_time,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AppUserAccessData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppUserAccessData struct{}"
	}

	return strings.Join([]string{"AppUserAccessData", string(data)}, " ")
}
