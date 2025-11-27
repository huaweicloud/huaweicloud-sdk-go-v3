package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterGroupResponse Response Object
type DeleteClusterGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClusterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterGroupResponse", string(data)}, " ")
}
