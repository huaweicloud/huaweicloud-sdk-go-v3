package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsfederationIdentityprovider
type OsfederationIdentityprovider struct {

	// 身份提供商ID。
	Id string `json:"id"`
}

func (o OsfederationIdentityprovider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsfederationIdentityprovider struct{}"
	}

	return strings.Join([]string{"OsfederationIdentityprovider", string(data)}, " ")
}
