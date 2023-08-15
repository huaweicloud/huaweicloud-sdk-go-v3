package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelBatchJobResponse Response Object
type CancelBatchJobResponse struct {

	// 取消成功，返回“deleted”。
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelBatchJobResponse struct{}"
	}

	return strings.Join([]string{"CancelBatchJobResponse", string(data)}, " ")
}
