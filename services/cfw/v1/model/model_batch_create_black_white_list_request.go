package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateBlackWhiteListRequest Request Object
type BatchCreateBlackWhiteListRequest struct {
	Body *AddBlackWhiteBatchDto `json:"body,omitempty"`
}

func (o BatchCreateBlackWhiteListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateBlackWhiteListRequest", string(data)}, " ")
}
