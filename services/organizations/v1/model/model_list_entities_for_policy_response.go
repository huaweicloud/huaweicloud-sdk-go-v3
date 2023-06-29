package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesForPolicyResponse Response Object
type ListEntitiesForPolicyResponse struct {

	// 结构列表，每个结构都包含有关指定策略附加到的实体的详细信息。
	AttachedEntities *[]EntityDto `json:"attached_entities,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEntitiesForPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesForPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListEntitiesForPolicyResponse", string(data)}, " ")
}
