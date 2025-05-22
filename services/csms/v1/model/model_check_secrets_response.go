package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSecretsResponse Response Object
type CheckSecretsResponse struct {

	// 凭据检测类型。
	Type *string `json:"type,omitempty"`

	// 凭据检测对应的响应数据体。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CheckSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSecretsResponse struct{}"
	}

	return strings.Join([]string{"CheckSecretsResponse", string(data)}, " ")
}
