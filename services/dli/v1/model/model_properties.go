package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Properties struct {
}

func (o Properties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Properties struct{}"
	}

	return strings.Join([]string{"Properties", string(data)}, " ")
}
