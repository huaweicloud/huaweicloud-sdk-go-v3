package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHbaConfsResponse Response Object
type DeleteHbaConfsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHbaConfsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHbaConfsResponse struct{}"
	}

	return strings.Join([]string{"DeleteHbaConfsResponse", string(data)}, " ")
}
