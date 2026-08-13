package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateResourceGroupRequestBody struct {
	Description string `json:"description"`
}

func (o UpdateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequestBody", string(data)}, " ")
}
