package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateCertsV2Request Request Object
type BatchAssociateCertsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id"`

	Body *AttachOrDetachCertsReqBody `json:"body,omitempty"`
}

func (o BatchAssociateCertsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateCertsV2Request struct{}"
	}

	return strings.Join([]string{"BatchAssociateCertsV2Request", string(data)}, " ")
}
