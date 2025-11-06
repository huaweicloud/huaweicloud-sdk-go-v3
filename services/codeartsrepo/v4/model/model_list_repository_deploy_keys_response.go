package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryDeployKeysResponse Response Object
type ListRepositoryDeployKeysResponse struct {

	// **参数解释：** 部署密钥列表。
	Body *[]DeployKeyDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryDeployKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryDeployKeysResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryDeployKeysResponse", string(data)}, " ")
}
