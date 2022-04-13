package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteConnctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConnctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnctionResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnctionResponse", string(data)}, " ")
}
