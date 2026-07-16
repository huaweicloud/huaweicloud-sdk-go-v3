package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPodIdentityAssociationsResponse Response Object
type ListPodIdentityAssociationsResponse struct {

	// **参数解释**: pod-identity 关联列表信息 **约束限制**： 不涉及
	Body           *[]PodIdentityAssociationResp `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPodIdentityAssociationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPodIdentityAssociationsResponse struct{}"
	}

	return strings.Join([]string{"ListPodIdentityAssociationsResponse", string(data)}, " ")
}
