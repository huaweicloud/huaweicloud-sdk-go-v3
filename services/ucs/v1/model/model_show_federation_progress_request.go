package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFederationProgressRequest Request Object
type ShowFederationProgressRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`
}

func (o ShowFederationProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFederationProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowFederationProgressRequest", string(data)}, " ")
}
