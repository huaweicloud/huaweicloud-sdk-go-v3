package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceShareResponse Response Object
type DeleteResourceShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceShareResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceShareResponse", string(data)}, " ")
}
