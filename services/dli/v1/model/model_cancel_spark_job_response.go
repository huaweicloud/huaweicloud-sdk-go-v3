package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSparkJobResponse Response Object
type CancelSparkJobResponse struct {

	// 取消成功，返回“deleted”。
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkJobResponse struct{}"
	}

	return strings.Join([]string{"CancelSparkJobResponse", string(data)}, " ")
}
