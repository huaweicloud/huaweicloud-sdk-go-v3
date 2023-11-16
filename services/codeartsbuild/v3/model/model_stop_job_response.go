package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobResponse Response Object
type StopJobResponse struct {

	// 构建停止结果
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
