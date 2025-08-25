package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeAccessKeysRequest Request Object
type AuthorizeAccessKeysRequest struct {

	// 所需要绑定应用的密码集群ID
	ClusterId string `json:"cluster_id"`

	Body *AuthorizeAccessKeysRequestBody `json:"body,omitempty"`
}

func (o AuthorizeAccessKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeAccessKeysRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeAccessKeysRequest", string(data)}, " ")
}
