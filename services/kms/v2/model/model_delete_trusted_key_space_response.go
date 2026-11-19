package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrustedKeySpaceResponse Response Object
type DeleteTrustedKeySpaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrustedKeySpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedKeySpaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrustedKeySpaceResponse", string(data)}, " ")
}
