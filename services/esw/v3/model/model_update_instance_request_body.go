package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceRequestBody struct {
	Instance *UpdateInstanceOption `json:"instance"`
}

func (o UpdateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequestBody", string(data)}, " ")
}
