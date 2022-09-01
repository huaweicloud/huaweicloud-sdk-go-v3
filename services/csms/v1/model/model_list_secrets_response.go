package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecretsResponse struct {

	// 凭据详情列表。
	Secrets *[]Secret `json:"secrets,omitempty" xml:"secrets"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretsResponse", string(data)}, " ")
}
