package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourceResponse struct{}"
	}

	return strings.Join([]string{"AddResourceResponse", string(data)}, " ")
}
