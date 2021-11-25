package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDedicatedHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateDedicatedHostResponse", string(data)}, " ")
}
