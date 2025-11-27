package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableFederationRequest Request Object
type EnableFederationRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	// 是否重试集群加入联邦
	Retryjoinall *bool `json:"retryjoinall,omitempty"`
}

func (o EnableFederationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableFederationRequest struct{}"
	}

	return strings.Join([]string{"EnableFederationRequest", string(data)}, " ")
}
