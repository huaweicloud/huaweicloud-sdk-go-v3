package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SaveMonitorItemConfigRequest struct {

	// 业务id
	XBusinessId int64 `json:"x-business-id" xml:"x-business-id"`

	Body *SaveMonitorItemParam `json:"body,omitempty" xml:"body"`
}

func (o SaveMonitorItemConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveMonitorItemConfigRequest struct{}"
	}

	return strings.Join([]string{"SaveMonitorItemConfigRequest", string(data)}, " ")
}
