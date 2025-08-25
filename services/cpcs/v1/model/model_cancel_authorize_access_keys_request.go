package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAuthorizeAccessKeysRequest Request Object
type CancelAuthorizeAccessKeysRequest struct {

	// 所需要解除授权的密码集群ID
	ClusterId string `json:"cluster_id"`

	Body *DeAuthorizeAccessKeysRequestBody `json:"body,omitempty"`
}

func (o CancelAuthorizeAccessKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAuthorizeAccessKeysRequest struct{}"
	}

	return strings.Join([]string{"CancelAuthorizeAccessKeysRequest", string(data)}, " ")
}
