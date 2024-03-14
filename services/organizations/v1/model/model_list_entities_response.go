package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesResponse Response Object
type ListEntitiesResponse struct {

	// 组织中的根、组织单元和账号的列表。
	Entities *[]EntityDto `json:"entities,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ListEntitiesResponse", string(data)}, " ")
}
