package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointUpdateReq struct {

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o EndpointUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointUpdateReq struct{}"
	}

	return strings.Join([]string{"EndpointUpdateReq", string(data)}, " ")
}
