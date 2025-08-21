package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectDeployKeysResponse Response Object
type ListProjectDeployKeysResponse struct {
	Body *[]DeployKeyDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectDeployKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDeployKeysResponse struct{}"
	}

	return strings.Join([]string{"ListProjectDeployKeysResponse", string(data)}, " ")
}
