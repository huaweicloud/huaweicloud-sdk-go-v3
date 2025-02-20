package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewResponseBody struct {
}

func (o BatchCreateApplicationViewResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewResponseBody struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewResponseBody", string(data)}, " ")
}
