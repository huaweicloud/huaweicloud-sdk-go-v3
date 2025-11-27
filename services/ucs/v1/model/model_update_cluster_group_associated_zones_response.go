package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedZonesResponse Response Object
type UpdateClusterGroupAssociatedZonesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterGroupAssociatedZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedZonesResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedZonesResponse", string(data)}, " ")
}
