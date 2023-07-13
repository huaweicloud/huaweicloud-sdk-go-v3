package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlanResponse Response Object
type CreatePlanResponse struct {

	// 接口调用成功返回的计划id
	PlanId *string `json:"plan_id,omitempty"`

	// 接口调用成功不返回，调用失败错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 接口调用成功不返回，调用失败错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlanResponse struct{}"
	}

	return strings.Join([]string{"CreatePlanResponse", string(data)}, " ")
}
