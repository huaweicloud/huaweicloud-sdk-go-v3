package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSshKeyResponse Response Object
type DeleteSshKeyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSshKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSshKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSshKeyResponse", string(data)}, " ")
}
