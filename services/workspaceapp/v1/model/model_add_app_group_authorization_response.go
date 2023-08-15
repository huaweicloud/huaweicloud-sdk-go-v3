package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppGroupAuthorizationResponse Response Object
type AddAppGroupAuthorizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAppGroupAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppGroupAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"AddAppGroupAuthorizationResponse", string(data)}, " ")
}
