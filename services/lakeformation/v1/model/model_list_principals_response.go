package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrincipalsResponse Response Object
type ListPrincipalsResponse struct {

	// 主体列表
	Principals *[]Principal `json:"principals,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrincipalsResponse struct{}"
	}

	return strings.Join([]string{"ListPrincipalsResponse", string(data)}, " ")
}
