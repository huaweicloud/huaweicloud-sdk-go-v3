package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhitelistResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistResponse", string(data)}, " ")
}
