package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsBucketInfoOwner 桶拥有者信息
type ObsBucketInfoOwner struct {

	// 显示名称
	DisplayName *string `json:"displayName,omitempty"`

	// 用户的DomainID，即帐号ID
	Id *string `json:"id,omitempty"`
}

func (o ObsBucketInfoOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucketInfoOwner struct{}"
	}

	return strings.Join([]string{"ObsBucketInfoOwner", string(data)}, " ")
}
