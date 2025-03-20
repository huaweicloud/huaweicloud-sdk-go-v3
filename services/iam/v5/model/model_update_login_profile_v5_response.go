package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginProfileV5Response Response Object
type UpdateLoginProfileV5Response struct {
	LoginProfile   *LoginProfile `json:"login_profile,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateLoginProfileV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProfileV5Response struct{}"
	}

	return strings.Join([]string{"UpdateLoginProfileV5Response", string(data)}, " ")
}
