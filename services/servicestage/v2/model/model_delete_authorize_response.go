package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthorizeResponse Response Object
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
