package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceAssociationsResponse Response Object
type ListResourceAssociationsResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**： 资源关联实例信息列表。 **取值范围**： 不涉及。
	ResourceAssociations *[]ResourceAssociation `json:"resource_associations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResourceAssociationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceAssociationsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceAssociationsResponse", string(data)}, " ")
}
