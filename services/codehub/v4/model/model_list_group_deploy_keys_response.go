package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupDeployKeysResponse Response Object
type ListGroupDeployKeysResponse struct {
	Body *[]DeployKeyDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupDeployKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupDeployKeysResponse struct{}"
	}

	return strings.Join([]string{"ListGroupDeployKeysResponse", string(data)}, " ")
}
