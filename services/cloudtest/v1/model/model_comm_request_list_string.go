package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestListString struct {
	Params []string `json:"params"`
}

func (o CommRequestListString) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestListString struct{}"
	}

	return strings.Join([]string{"CommRequestListString", string(data)}, " ")
}
