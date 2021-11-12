package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceRequestBody struct {
	Resize *ResizeInstanceOption `json:"resize"`
}

func (o ResizeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequestBody", string(data)}, " ")
}
