package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockPortResponseModel 通道号解绑服务号返回体。
type UnlockPortResponseModel struct {

	// 信息。
	Message *string `json:"message,omitempty"`
}

func (o UnlockPortResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockPortResponseModel struct{}"
	}

	return strings.Join([]string{"UnlockPortResponseModel", string(data)}, " ")
}
