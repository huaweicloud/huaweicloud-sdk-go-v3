package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSourceResponse", string(data)}, " ")
}
