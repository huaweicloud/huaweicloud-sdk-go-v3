package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ReclaimPartnerCouponsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ReclaimPartnerCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimPartnerCouponsResponse struct{}"
	}

	return strings.Join([]string{"ReclaimPartnerCouponsResponse", string(data)}, " ")
}
