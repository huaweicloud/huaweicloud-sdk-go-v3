package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSyncRequest 取消IAM同步请求的指定参数
type CancelSyncRequest struct {

	// 指定取消同步的IAM用户组。
	GroupNames *[]string `json:"group_names,omitempty"`

	// 指定同步的IAM用户。
	UserNames *[]string `json:"user_names,omitempty"`
}

func (o CancelSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSyncRequest struct{}"
	}

	return strings.Join([]string{"CancelSyncRequest", string(data)}, " ")
}
