package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletefavoriteResponse Response Object
type DeletefavoriteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletefavoriteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletefavoriteResponse struct{}"
	}

	return strings.Join([]string{"DeletefavoriteResponse", string(data)}, " ")
}
