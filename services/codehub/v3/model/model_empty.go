package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Empty struct {
}

func (o Empty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Empty struct{}"
	}

	return strings.Join([]string{"Empty", string(data)}, " ")
}
