package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type NovaCreateServersRequestBody struct {
	Server *NovaCreateServersOption `json:"server" xml:"server"`

	OsschedulerHints *NovaCreateServersSchedulerHint `json:"os:scheduler_hints,omitempty" xml:"os:scheduler_hints"`
}

func (o NovaCreateServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersRequestBody struct{}"
	}

	return strings.Join([]string{"NovaCreateServersRequestBody", string(data)}, " ")
}
