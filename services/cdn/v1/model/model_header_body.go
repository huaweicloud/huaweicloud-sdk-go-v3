package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HeaderBody struct {
	Headers *HeaderMap `json:"headers,omitempty"`
}

func (o HeaderBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeaderBody struct{}"
	}

	return strings.Join([]string{"HeaderBody", string(data)}, " ")
}
