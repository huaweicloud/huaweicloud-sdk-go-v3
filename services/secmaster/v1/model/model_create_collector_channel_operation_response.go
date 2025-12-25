package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelOperationResponse Response Object
type CreateCollectorChannelOperationResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCollectorChannelOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelOperationResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelOperationResponse", string(data)}, " ")
}
