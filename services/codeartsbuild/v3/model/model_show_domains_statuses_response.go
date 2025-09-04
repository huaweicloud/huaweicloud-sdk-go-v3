package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsStatusesResponse Response Object
type ShowDomainsStatusesResponse struct {

	// 结果
	Result *[]DomainStatuses `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainsStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsStatusesResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainsStatusesResponse", string(data)}, " ")
}
