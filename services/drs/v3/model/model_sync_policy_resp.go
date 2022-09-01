package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncPolicyResp struct {

	// 任务ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 状态 - success：成功 - failed：失败
	Status *string `json:"status,omitempty" xml:"status"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o SyncPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyResp struct{}"
	}

	return strings.Join([]string{"SyncPolicyResp", string(data)}, " ")
}
