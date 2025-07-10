package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTheJobResponse Response Object
type StopTheJobResponse struct {

	// 是否停止成功
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopTheJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTheJobResponse struct{}"
	}

	return strings.Join([]string{"StopTheJobResponse", string(data)}, " ")
}
