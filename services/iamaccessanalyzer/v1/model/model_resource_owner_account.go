package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceOwnerAccount 拥有资源的账号ID。
type ResourceOwnerAccount struct {
}

func (o ResourceOwnerAccount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceOwnerAccount struct{}"
	}

	return strings.Join([]string{"ResourceOwnerAccount", string(data)}, " ")
}
