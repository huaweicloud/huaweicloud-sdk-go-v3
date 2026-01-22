package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateBlackWhiteListResponse Response Object
type BatchCreateBlackWhiteListResponse struct {
	Data           *AddBlackWhiteListBatchVo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchCreateBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateBlackWhiteListResponse", string(data)}, " ")
}
