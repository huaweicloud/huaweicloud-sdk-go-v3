package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SuccessResponse struct {
}

func (o SuccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SuccessResponse struct{}"
	}

	return strings.Join([]string{"SuccessResponse", string(data)}, " ")
}
