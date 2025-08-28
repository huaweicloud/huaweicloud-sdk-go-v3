package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRegistryResponse Response Object
type AddRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRegistryResponse struct{}"
	}

	return strings.Join([]string{"AddRegistryResponse", string(data)}, " ")
}
