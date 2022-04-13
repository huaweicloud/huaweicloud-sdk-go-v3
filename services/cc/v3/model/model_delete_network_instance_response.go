package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteNetworkInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteNetworkInstanceResponse", string(data)}, " ")
}
