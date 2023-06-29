package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockPortResponseModel 通道号绑定服务号返回返回体。
type LockPortResponseModel struct {

	// 信息。
	Message *string `json:"message,omitempty"`
}

func (o LockPortResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockPortResponseModel struct{}"
	}

	return strings.Join([]string{"LockPortResponseModel", string(data)}, " ")
}
