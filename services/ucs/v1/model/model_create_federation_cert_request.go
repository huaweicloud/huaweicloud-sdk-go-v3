package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationCertRequest Request Object
type CreateFederationCertRequest struct {

	// 舰队ID
	Clustergroupid string `json:"clustergroupid"`

	Body *CreateFederationCertRequestBody `json:"body,omitempty"`
}

func (o CreateFederationCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationCertRequest struct{}"
	}

	return strings.Join([]string{"CreateFederationCertRequest", string(data)}, " ")
}
