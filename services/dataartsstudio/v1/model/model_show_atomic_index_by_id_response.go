package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAtomicIndexByIdResponse Response Object
type ShowAtomicIndexByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAtomicIndexByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAtomicIndexByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAtomicIndexByIdResponse", string(data)}, " ")
}
