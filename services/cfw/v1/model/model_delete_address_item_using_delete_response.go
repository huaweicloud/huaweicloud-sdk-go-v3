package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressItemUsingDeleteResponse Response Object
type DeleteAddressItemUsingDeleteResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteAddressItemUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemUsingDeleteResponse", string(data)}, " ")
}
