package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAddressSetListUsingGetResponse struct {
	Data           *AddressSetListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListAddressSetListUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetListUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListAddressSetListUsingGetResponse", string(data)}, " ")
}
