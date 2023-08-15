package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePwdResponse Response Object
type UpdatePwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePwdResponse struct{}"
	}

	return strings.Join([]string{"UpdatePwdResponse", string(data)}, " ")
}
