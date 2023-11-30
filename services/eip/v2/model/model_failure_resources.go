package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailureResources 带宽失败成功对象
type FailureResources struct {

	// - 功能说明：更新失败的带宽id
	Id *string `json:"id,omitempty"`

	// - 功能说明：错误码
	Code *string `json:"code,omitempty"`

	// - 功能说明：错误信息
	Message *string `json:"message,omitempty"`
}

func (o FailureResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailureResources struct{}"
	}

	return strings.Join([]string{"FailureResources", string(data)}, " ")
}
