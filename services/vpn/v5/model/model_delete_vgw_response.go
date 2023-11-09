package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVgwResponse Response Object
type DeleteVgwResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVgwResponse struct{}"
	}

	return strings.Join([]string{"DeleteVgwResponse", string(data)}, " ")
}
