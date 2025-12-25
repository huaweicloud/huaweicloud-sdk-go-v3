package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationDto struct {
	Action *ChannelAction `json:"action,omitempty"`
}

func (o OperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationDto struct{}"
	}

	return strings.Join([]string{"OperationDto", string(data)}, " ")
}
