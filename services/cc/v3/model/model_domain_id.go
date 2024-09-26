package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainId 实例所属账号ID。
type DomainId struct {

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`
}

func (o DomainId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainId struct{}"
	}

	return strings.Join([]string{"DomainId", string(data)}, " ")
}
