package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqAssociatedResourceOpenStatus 开通或关闭关联资源标签继承能力
type ReqAssociatedResourceOpenStatus struct {

	// 状态
	Status string `json:"status"`
}

func (o ReqAssociatedResourceOpenStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqAssociatedResourceOpenStatus struct{}"
	}

	return strings.Join([]string{"ReqAssociatedResourceOpenStatus", string(data)}, " ")
}
