package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareAssociationsResponse Response Object
type SearchResourceShareAssociationsResponse struct {

	// 绑定的详细信息列表。
	ResourceShareAssociations *[]ResourceShareAssociation `json:"resource_share_associations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchResourceShareAssociationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareAssociationsResponse struct{}"
	}

	return strings.Join([]string{"SearchResourceShareAssociationsResponse", string(data)}, " ")
}
