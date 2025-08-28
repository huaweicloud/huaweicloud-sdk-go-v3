package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRegistryResponse Response Object
type DeleteRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRegistryResponse struct{}"
	}

	return strings.Join([]string{"DeleteRegistryResponse", string(data)}, " ")
}
