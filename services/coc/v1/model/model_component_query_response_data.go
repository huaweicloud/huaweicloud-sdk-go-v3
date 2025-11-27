package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentQueryResponseData 描述。
type ComponentQueryResponseData struct {

	// CMDB分配的uuid。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// code。
	Code *string `json:"code,omitempty"`

	// 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// 应用id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 企业项目id。
	EpId *string `json:"ep_id,omitempty"`
}

func (o ComponentQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentQueryResponseData struct{}"
	}

	return strings.Join([]string{"ComponentQueryResponseData", string(data)}, " ")
}
