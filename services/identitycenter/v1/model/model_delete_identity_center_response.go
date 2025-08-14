package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIdentityCenterResponse Response Object
type DeleteIdentityCenterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIdentityCenterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIdentityCenterResponse struct{}"
	}

	return strings.Join([]string{"DeleteIdentityCenterResponse", string(data)}, " ")
}
