package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveMonitorItemConfigRequest Request Object
type SaveMonitorItemConfigRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *SaveMonitorItemParam `json:"body,omitempty"`
}

func (o SaveMonitorItemConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveMonitorItemConfigRequest struct{}"
	}

	return strings.Join([]string{"SaveMonitorItemConfigRequest", string(data)}, " ")
}
