package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopPoolResponse Response Object
type ResizeDesktopPoolResponse struct {

	// 创建云桌面总任务id
	JobId *string `json:"job_id,omitempty"`

	// 变更订单错误处理策略。cbc调用返回值。设置为 NO_WORKORDER。云运营平台会认为无法开通成功，退费给客户后，不会再发运维工单给云服务，而由云服务自己去闭环处理对应问题。
	ErrorPolicy *string `json:"error_policy,omitempty"`

	// 按需桌面变更规格返回的任务信息。
	Jobs           *[]ResizeDesktopPoolJobResponse `json:"jobs,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ResizeDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolResponse", string(data)}, " ")
}
