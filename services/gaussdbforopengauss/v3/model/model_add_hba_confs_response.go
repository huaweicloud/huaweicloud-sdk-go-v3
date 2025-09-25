package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddHbaConfsResponse Response Object
type AddHbaConfsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddHbaConfsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHbaConfsResponse struct{}"
	}

	return strings.Join([]string{"AddHbaConfsResponse", string(data)}, " ")
}
