package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateipRequestBody
type CreatePrivateipRequestBody struct {

	// 私有IP列表对象
	Privateips []CreatePrivateipOption `json:"privateips"`
}

func (o CreatePrivateipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateipRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateipRequestBody", string(data)}, " ")
}
