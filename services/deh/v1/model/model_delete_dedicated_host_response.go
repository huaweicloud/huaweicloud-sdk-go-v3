package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDedicatedHostResponse Response Object
type DeleteDedicatedHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteDedicatedHostResponse", string(data)}, " ")
}
