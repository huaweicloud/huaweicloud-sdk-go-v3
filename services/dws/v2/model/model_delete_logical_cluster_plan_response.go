package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogicalClusterPlanResponse Response Object
type DeleteLogicalClusterPlanResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogicalClusterPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogicalClusterPlanResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogicalClusterPlanResponse", string(data)}, " ")
}
