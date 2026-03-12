package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesForDryRunPolicyResponse Response Object
type ListEntitiesForDryRunPolicyResponse struct {

	// 结构列表，每个结构都包含有关指定策略附加到的实体的详细信息。
	AttachedEntities *[]EntityDto `json:"attached_entities,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEntitiesForDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesForDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListEntitiesForDryRunPolicyResponse", string(data)}, " ")
}
