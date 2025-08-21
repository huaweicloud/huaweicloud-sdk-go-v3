package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveOrderRequest Request Object
type SaveOrderRequest struct {
	Body *SaveOrderBody `json:"body,omitempty"`
}

func (o SaveOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveOrderRequest struct{}"
	}

	return strings.Join([]string{"SaveOrderRequest", string(data)}, " ")
}
