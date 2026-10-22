package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2RayFederationsResponse Response Object
type ListV2RayFederationsResponse struct {

	// **参数解释**：Ray联邦端点总记录数。 **取值范围**：大于等于0的整数。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：Ray联邦端点列表。
	RayFederations *[]V2RayFederationInfo `json:"ray_federations,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListV2RayFederationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2RayFederationsResponse struct{}"
	}

	return strings.Join([]string{"ListV2RayFederationsResponse", string(data)}, " ")
}
