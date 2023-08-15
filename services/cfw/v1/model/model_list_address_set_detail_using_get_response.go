package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressSetDetailUsingGetResponse Response Object
type ListAddressSetDetailUsingGetResponse struct {
	Data           *AddressSetDetailResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListAddressSetDetailUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetDetailUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListAddressSetDetailUsingGetResponse", string(data)}, " ")
}
