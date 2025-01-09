package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAuthorizedMailResponse Response Object
type SendAuthorizedMailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendAuthorizedMailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAuthorizedMailResponse struct{}"
	}

	return strings.Join([]string{"SendAuthorizedMailResponse", string(data)}, " ")
}
