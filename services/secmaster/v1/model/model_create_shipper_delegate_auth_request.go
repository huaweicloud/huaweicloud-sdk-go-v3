package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShipperDelegateAuthRequest Request Object
type CreateShipperDelegateAuthRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *CreateShipperDelegateAuthRequestBody `json:"body,omitempty"`
}

func (o CreateShipperDelegateAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperDelegateAuthRequest struct{}"
	}

	return strings.Join([]string{"CreateShipperDelegateAuthRequest", string(data)}, " ")
}
