package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnLogConfigRequestBodyContent struct {

	// 日志组ID
	LogGroupId string `json:"log_group_id"`

	// 日志流ID
	LogStreamId string `json:"log_stream_id"`
}

func (o UpdateVpnLogConfigRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnLogConfigRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnLogConfigRequestBodyContent", string(data)}, " ")
}
