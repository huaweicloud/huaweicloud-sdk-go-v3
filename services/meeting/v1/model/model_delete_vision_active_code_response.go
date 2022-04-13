package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVisionActiveCodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVisionActiveCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVisionActiveCodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteVisionActiveCodeResponse", string(data)}, " ")
}
