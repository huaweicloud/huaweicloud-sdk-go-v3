package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoginProfileV5Response Response Object
type ShowLoginProfileV5Response struct {
	LoginProfile   *LoginProfile `json:"login_profile,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowLoginProfileV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoginProfileV5Response struct{}"
	}

	return strings.Join([]string{"ShowLoginProfileV5Response", string(data)}, " ")
}
