package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAcceleratorResponse Response Object
type DeleteAcceleratorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAcceleratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAcceleratorResponse struct{}"
	}

	return strings.Join([]string{"DeleteAcceleratorResponse", string(data)}, " ")
}
