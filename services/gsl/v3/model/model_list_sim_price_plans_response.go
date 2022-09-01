package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSimPricePlansResponse struct {
	Body           *[]SimPricePlanVo `json:"body,omitempty" xml:"body"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSimPricePlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimPricePlansResponse struct{}"
	}

	return strings.Join([]string{"ListSimPricePlansResponse", string(data)}, " ")
}
