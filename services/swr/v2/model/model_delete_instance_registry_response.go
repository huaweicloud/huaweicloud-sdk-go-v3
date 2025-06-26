package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceRegistryResponse Response Object
type DeleteInstanceRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRegistryResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRegistryResponse", string(data)}, " ")
}
