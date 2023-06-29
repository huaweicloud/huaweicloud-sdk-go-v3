package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressSetInfoUsingDeleteResponse Response Object
type DeleteAddressSetInfoUsingDeleteResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteAddressSetInfoUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressSetInfoUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressSetInfoUsingDeleteResponse", string(data)}, " ")
}
