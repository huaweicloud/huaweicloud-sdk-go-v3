package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTsviRequest Request Object
type BatchUpdateTsviRequest struct {
	Body *UpdateTsviReq `json:"body,omitempty"`
}

func (o BatchUpdateTsviRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTsviRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateTsviRequest", string(data)}, " ")
}
