package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableFederationRequest Request Object
type DisableFederationRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`
}

func (o DisableFederationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableFederationRequest struct{}"
	}

	return strings.Join([]string{"DisableFederationRequest", string(data)}, " ")
}
