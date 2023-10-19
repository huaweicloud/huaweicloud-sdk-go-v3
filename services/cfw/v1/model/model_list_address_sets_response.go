package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressSetsResponse Response Object
type ListAddressSetsResponse struct {
	Data           *AddressSetListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListAddressSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetsResponse struct{}"
	}

	return strings.Join([]string{"ListAddressSetsResponse", string(data)}, " ")
}
