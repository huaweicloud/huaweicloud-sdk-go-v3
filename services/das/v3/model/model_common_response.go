package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponse struct {
}

func (o CommonResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponse struct{}"
	}

	return strings.Join([]string{"CommonResponse", string(data)}, " ")
}
