package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateDomainV2Request Request Object
type AssociateDomainV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	Body *UrlDomainCreate `json:"body,omitempty"`
}

func (o AssociateDomainV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDomainV2Request struct{}"
	}

	return strings.Join([]string{"AssociateDomainV2Request", string(data)}, " ")
}
