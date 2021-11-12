package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainTrafficDetailResponse struct {
	// 采样数据列表。

	DataList *[]TrafficData `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainTrafficDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainTrafficDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDomainTrafficDetailResponse", string(data)}, " ")
}
