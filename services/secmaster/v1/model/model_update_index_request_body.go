package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIndexRequestBody struct {
}

func (o UpdateIndexRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndexRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIndexRequestBody", string(data)}, " ")
}
