package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateHealthmonitorRequestBody struct {
	Healthmonitor *CreateHealthmonitorReq `json:"healthmonitor"`
}

func (o CreateHealthmonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequestBody", string(data)}, " ")
}
