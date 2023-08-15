package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type L2Id struct {
}

func (o L2Id) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L2Id struct{}"
	}

	return strings.Join([]string{"L2Id", string(data)}, " ")
}
