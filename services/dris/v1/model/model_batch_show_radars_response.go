package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowRadarsResponse Response Object
type BatchShowRadarsResponse struct {

	// **参数说明**：总数
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：数据列表
	Radars         *[]RadarResourceDto `json:"radars,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchShowRadarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowRadarsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowRadarsResponse", string(data)}, " ")
}
