package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootServersRequestBody This is a auto create Body Object
type BatchRebootServersRequestBody struct {
	Reboot *BatchRebootSeversOption `json:"reboot"`
}

func (o BatchRebootServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRebootServersRequestBody", string(data)}, " ")
}
