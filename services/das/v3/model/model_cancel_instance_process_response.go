package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceProcessResponse Response Object
type CancelInstanceProcessResponse struct {

	// 进程数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CancelInstanceProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceProcessResponse struct{}"
	}

	return strings.Join([]string{"CancelInstanceProcessResponse", string(data)}, " ")
}
