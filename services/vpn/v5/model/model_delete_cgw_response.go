package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCgwResponse Response Object
type DeleteCgwResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCgwResponse struct{}"
	}

	return strings.Join([]string{"DeleteCgwResponse", string(data)}, " ")
}
