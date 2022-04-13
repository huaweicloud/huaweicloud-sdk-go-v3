package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type LiveResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveResponse struct{}"
	}

	return strings.Join([]string{"LiveResponse", string(data)}, " ")
}
