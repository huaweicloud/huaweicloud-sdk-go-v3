package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressSetDetailResponse Response Object
type ListAddressSetDetailResponse struct {
	Data           *AddressSetDetailResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListAddressSetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAddressSetDetailResponse", string(data)}, " ")
}
