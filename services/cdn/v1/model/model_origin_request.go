package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OriginRequest struct {
	Origin *ResourceBody `json:"origin"`
}

func (o OriginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginRequest struct{}"
	}

	return strings.Join([]string{"OriginRequest", string(data)}, " ")
}
