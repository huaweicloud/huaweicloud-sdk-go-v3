package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeFederationRequest Request Object
type UpgradeFederationRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *UpgradeFederationRequestBody `json:"body,omitempty"`
}

func (o UpgradeFederationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeFederationRequest struct{}"
	}

	return strings.Join([]string{"UpgradeFederationRequest", string(data)}, " ")
}
