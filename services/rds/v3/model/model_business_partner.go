package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessPartner 服务商信息。
type BusinessPartner struct {

	// 服务商ID。
	BpDomainId *string `json:"bp_domain_id,omitempty"`

	// 服务商名称。
	BpName *string `json:"bp_name,omitempty"`

	// 优先级，整数取值范围1-100，数值越小优先级越高。
	Order *int32 `json:"order,omitempty"`

	// 是否是国际站服务商。
	International *bool `json:"international,omitempty"`
}

func (o BusinessPartner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessPartner struct{}"
	}

	return strings.Join([]string{"BusinessPartner", string(data)}, " ")
}
