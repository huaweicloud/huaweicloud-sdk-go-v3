package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVaultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultResponse", string(data)}, " ")
}
