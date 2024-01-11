package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadQueueUserReqUserList struct {

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o WorkloadQueueUserReqUserList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueUserReqUserList struct{}"
	}

	return strings.Join([]string{"WorkloadQueueUserReqUserList", string(data)}, " ")
}
