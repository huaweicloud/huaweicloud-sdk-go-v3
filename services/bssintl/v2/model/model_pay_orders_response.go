package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayOrdersResponse Response Object
type PayOrdersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PayOrdersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayOrdersResponse struct{}"
	}

	return strings.Join([]string{"PayOrdersResponse", string(data)}, " ")
}
