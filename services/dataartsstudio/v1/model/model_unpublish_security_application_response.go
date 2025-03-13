package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishSecurityApplicationResponse Response Object
type UnpublishSecurityApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnpublishSecurityApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishSecurityApplicationResponse struct{}"
	}

	return strings.Join([]string{"UnpublishSecurityApplicationResponse", string(data)}, " ")
}
