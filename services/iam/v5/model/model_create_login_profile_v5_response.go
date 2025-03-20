package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoginProfileV5Response Response Object
type CreateLoginProfileV5Response struct {
	LoginProfile   *LoginProfile `json:"login_profile,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateLoginProfileV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginProfileV5Response struct{}"
	}

	return strings.Join([]string{"CreateLoginProfileV5Response", string(data)}, " ")
}
