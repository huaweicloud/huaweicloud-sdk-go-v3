package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogicalClusterPlanResponse Response Object
type EnableLogicalClusterPlanResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableLogicalClusterPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogicalClusterPlanResponse struct{}"
	}

	return strings.Join([]string{"EnableLogicalClusterPlanResponse", string(data)}, " ")
}
