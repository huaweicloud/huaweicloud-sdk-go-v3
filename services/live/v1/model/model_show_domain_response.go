package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainResponse struct {

	// 查询结果的总数量
	Total float32 `json:"total,omitempty" xml:"total"`

	// 直播域名列表
	DomainInfo     *[]DecoupledLiveDomainInfo `json:"domain_info,omitempty" xml:"domain_info"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainResponse", string(data)}, " ")
}
