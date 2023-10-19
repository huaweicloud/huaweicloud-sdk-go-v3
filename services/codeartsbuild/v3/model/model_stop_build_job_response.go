package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBuildJobResponse Response Object
type StopBuildJobResponse struct {

	// 是否停止成功
	Success *bool `json:"success,omitempty"`

	// 返回结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBuildJobResponse struct{}"
	}

	return strings.Join([]string{"StopBuildJobResponse", string(data)}, " ")
}
