package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerConsoleOutputRequest Request Object
type ShowServerConsoleOutputRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	// - 参数解释： 请求log行数。 - 约束限制： 不涉及。 - 取值范围： 大于等于-1。其中-1代表不限长度输出。 - 默认取值： 不填时默认50。
	Length *int32 `json:"length,omitempty"`
}

func (o ShowServerConsoleOutputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerConsoleOutputRequest struct{}"
	}

	return strings.Join([]string{"ShowServerConsoleOutputRequest", string(data)}, " ")
}
