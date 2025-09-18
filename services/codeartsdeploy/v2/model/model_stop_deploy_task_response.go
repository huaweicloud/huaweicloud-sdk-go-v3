package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopDeployTaskResponse Response Object
type StopDeployTaskResponse struct {
	Result *CancelInfo `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopDeployTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDeployTaskResponse struct{}"
	}

	return strings.Join([]string{"StopDeployTaskResponse", string(data)}, " ")
}
