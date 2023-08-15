package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWebinarResponse Response Object
type DeleteWebinarResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWebinarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebinarResponse struct{}"
	}

	return strings.Join([]string{"DeleteWebinarResponse", string(data)}, " ")
}
