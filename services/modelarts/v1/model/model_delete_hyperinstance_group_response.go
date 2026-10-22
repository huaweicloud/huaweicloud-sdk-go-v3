package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperinstanceGroupResponse Response Object
type DeleteHyperinstanceGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHyperinstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperinstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteHyperinstanceGroupResponse", string(data)}, " ")
}
