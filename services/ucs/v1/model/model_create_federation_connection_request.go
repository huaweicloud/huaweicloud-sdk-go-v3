package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationConnectionRequest Request Object
type CreateFederationConnectionRequest struct {

	// 舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *CreateFederationConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateFederationConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateFederationConnectionRequest", string(data)}, " ")
}
