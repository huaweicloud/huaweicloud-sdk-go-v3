package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVaultTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagResponse", string(data)}, " ")
}
