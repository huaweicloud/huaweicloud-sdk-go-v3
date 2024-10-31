package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseDetailResponse Response Object
type ListDomainParseDetailResponse struct {

	// 域名解析ip列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDomainParseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDomainParseDetailResponse", string(data)}, " ")
}
