package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetailResponse Response Object
type ShowDomainDetailResponse struct {

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// cname
	Cname *string `json:"cname,omitempty"`

	// 域名状态 0-正常，1-冻结
	DomainStatus *string `json:"domain_status,omitempty"`

	// 是否开启pp 0-关闭，1-开启
	PpEnable       *int32 `json:"pp_enable,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDomainDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailResponse", string(data)}, " ")
}
