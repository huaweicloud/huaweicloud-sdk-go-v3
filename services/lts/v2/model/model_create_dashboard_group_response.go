package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashboardGroupResponse Response Object
type CreateDashboardGroupResponse struct {

	// 响应结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDashboardGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateDashboardGroupResponse", string(data)}, " ")
}
