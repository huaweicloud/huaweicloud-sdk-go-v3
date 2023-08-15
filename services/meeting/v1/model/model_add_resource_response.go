package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddResourceResponse Response Object
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
