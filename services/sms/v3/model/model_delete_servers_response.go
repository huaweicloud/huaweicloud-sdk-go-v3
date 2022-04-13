package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteServersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersResponse struct{}"
	}

	return strings.Join([]string{"DeleteServersResponse", string(data)}, " ")
}
