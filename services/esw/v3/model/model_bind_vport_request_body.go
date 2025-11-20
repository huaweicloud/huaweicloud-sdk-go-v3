package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindVportRequestBody struct {
	Vport *Vport `json:"vport"`
}

func (o BindVportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindVportRequestBody struct{}"
	}

	return strings.Join([]string{"BindVportRequestBody", string(data)}, " ")
}
