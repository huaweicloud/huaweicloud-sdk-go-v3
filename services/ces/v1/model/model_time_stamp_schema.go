package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeStampSchema struct {
}

func (o TimeStampSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeStampSchema struct{}"
	}

	return strings.Join([]string{"TimeStampSchema", string(data)}, " ")
}
