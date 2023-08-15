package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlanListResponse Response Object
type ShowPlanListResponse struct {

	// 项目下查询测试计划列表返回结构
	Body           *[]TestPlanDetail `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPlanListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanListResponse struct{}"
	}

	return strings.Join([]string{"ShowPlanListResponse", string(data)}, " ")
}
