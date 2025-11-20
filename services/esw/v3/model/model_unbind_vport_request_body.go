package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindVportRequestBody struct {
	Vport *Vport `json:"vport"`
}

func (o UnbindVportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindVportRequestBody struct{}"
	}

	return strings.Join([]string{"UnbindVportRequestBody", string(data)}, " ")
}
