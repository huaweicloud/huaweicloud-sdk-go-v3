package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FederationVersionSpec struct {

	// 联邦版本
	FederationVersion *string `json:"FederationVersion,omitempty"`
}

func (o FederationVersionSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FederationVersionSpec struct{}"
	}

	return strings.Join([]string{"FederationVersionSpec", string(data)}, " ")
}
