package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentResponse Response Object
type DeleteDeploymentResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentResponse", string(data)}, " ")
}
