package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationConnectionResponse Response Object
type CreateFederationConnectionResponse struct {

	// vpcep终端节点的id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFederationConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateFederationConnectionResponse", string(data)}, " ")
}
