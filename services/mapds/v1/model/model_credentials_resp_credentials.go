package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CredentialsRespCredentials struct {

	// 凭证id
	Uuid *string `json:"uuid,omitempty"`

	// 凭证value
	Key *string `json:"key,omitempty"`

	// 凭证创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 凭证描述
	Description *string `json:"description,omitempty"`

	// 凭证状态
	Status *string `json:"status,omitempty"`
}

func (o CredentialsRespCredentials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CredentialsRespCredentials struct{}"
	}

	return strings.Join([]string{"CredentialsRespCredentials", string(data)}, " ")
}
