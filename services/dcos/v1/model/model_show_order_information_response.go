package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderInformationResponse Response Object
type ShowOrderInformationResponse struct {

	// 联系号码
	Phone *[]UserPhone `json:"phone,omitempty"`

	// sop信息
	Sop            *[]DefaultSop `json:"sop,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowOrderInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderInformationResponse struct{}"
	}

	return strings.Join([]string{"ShowOrderInformationResponse", string(data)}, " ")
}
