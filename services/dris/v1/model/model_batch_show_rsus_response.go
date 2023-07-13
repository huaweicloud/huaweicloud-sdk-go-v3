package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowRsusResponse Response Object
type BatchShowRsusResponse struct {

	// **参数说明**：返回RSU的总体数量。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：RSU数据列表。
	Rsus           *[]RsuDto `json:"rsus,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchShowRsusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowRsusResponse struct{}"
	}

	return strings.Join([]string{"BatchShowRsusResponse", string(data)}, " ")
}
