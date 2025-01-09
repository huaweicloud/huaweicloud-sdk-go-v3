package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendEmailResponse Response Object
type SendEmailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendEmailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendEmailResponse struct{}"
	}

	return strings.Join([]string{"SendEmailResponse", string(data)}, " ")
}
