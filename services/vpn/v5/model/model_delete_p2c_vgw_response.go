package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVgwResponse Response Object
type DeleteP2cVgwResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteP2cVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVgwResponse struct{}"
	}

	return strings.Join([]string{"DeleteP2cVgwResponse", string(data)}, " ")
}
