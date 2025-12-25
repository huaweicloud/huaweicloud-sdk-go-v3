package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteShipperRequestBody struct {

	// 投递ID
	Ids []string `json:"ids"`
}

func (o DeleteShipperRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShipperRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteShipperRequestBody", string(data)}, " ")
}
