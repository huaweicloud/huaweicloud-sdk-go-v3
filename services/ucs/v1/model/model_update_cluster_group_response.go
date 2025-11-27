package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupResponse Response Object
type UpdateClusterGroupResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupResponse", string(data)}, " ")
}
