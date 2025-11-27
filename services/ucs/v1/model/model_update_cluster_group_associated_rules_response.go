package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedRulesResponse Response Object
type UpdateClusterGroupAssociatedRulesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterGroupAssociatedRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedRulesResponse", string(data)}, " ")
}
