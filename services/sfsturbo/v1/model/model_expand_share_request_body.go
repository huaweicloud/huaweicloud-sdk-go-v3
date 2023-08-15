package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandShareRequestBody struct {
	Extend *Extend `json:"extend"`
}

func (o ExpandShareRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandShareRequestBody", string(data)}, " ")
}
