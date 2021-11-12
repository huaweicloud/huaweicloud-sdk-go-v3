package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAuthorizeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAuthorizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorizeResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizeResponse", string(data)}, " ")
}
