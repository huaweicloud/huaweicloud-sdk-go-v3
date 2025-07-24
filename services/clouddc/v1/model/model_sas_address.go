package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SasAddress SAS地址
type SasAddress struct {
}

func (o SasAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SasAddress struct{}"
	}

	return strings.Join([]string{"SasAddress", string(data)}, " ")
}
