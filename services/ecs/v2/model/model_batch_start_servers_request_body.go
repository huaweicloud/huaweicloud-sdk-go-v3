package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartServersRequestBody This is a auto create Body Object
type BatchStartServersRequestBody struct {
	OsStart *BatchStartServersOption `json:"os-start"`
}

func (o BatchStartServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStartServersRequestBody", string(data)}, " ")
}
