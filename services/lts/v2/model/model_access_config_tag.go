package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessConfigTag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o AccessConfigTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigTag struct{}"
	}

	return strings.Join([]string{"AccessConfigTag", string(data)}, " ")
}
