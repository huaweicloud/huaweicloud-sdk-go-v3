package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapId CapId，32~36位的英文、数字、短横组合
type CapId struct {
}

func (o CapId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapId struct{}"
	}

	return strings.Join([]string{"CapId", string(data)}, " ")
}
