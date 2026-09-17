package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryBinlogPartResponse Response Object
type RetryBinlogPartResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o RetryBinlogPartResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBinlogPartResponse struct{}"
	}

	return strings.Join([]string{"RetryBinlogPartResponse", string(data)}, " ")
}
