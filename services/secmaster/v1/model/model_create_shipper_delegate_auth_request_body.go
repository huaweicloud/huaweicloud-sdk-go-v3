package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateShipperDelegateAuthRequestBody struct {

	// 委托的名称。
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o CreateShipperDelegateAuthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperDelegateAuthRequestBody struct{}"
	}

	return strings.Join([]string{"CreateShipperDelegateAuthRequestBody", string(data)}, " ")
}
