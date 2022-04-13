package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstanceRequestBody struct {
	Source *Source `json:"source"`

	Target *Target `json:"target"`
}

func (o RestoreInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequestBody", string(data)}, " ")
}
