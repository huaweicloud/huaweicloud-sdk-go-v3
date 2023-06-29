package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCredentialResponse Response Object
type CreateCredentialResponse struct {

	// 凭证
	Key *string `json:"key,omitempty"`

	// 创建凭证的时间UTC时间格式：YYYY-mm-dd'T'HH:mm:ss.SSSSSS'Z'，e.g. \"2020-01-08T06:26:08.123059Z\"
	CreateTime *string `json:"create_time,omitempty"`

	// 凭证的描述信息。
	Description *string `json:"description,omitempty"`

	// 凭证状态“ACTIVE”
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCredentialResponse struct{}"
	}

	return strings.Join([]string{"CreateCredentialResponse", string(data)}, " ")
}
