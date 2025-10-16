package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthObjectScope 鉴权范围
type AuthObjectScope struct {
	Cluster *AuthObjectScopeCluster `json:"cluster,omitempty"`
}

func (o AuthObjectScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthObjectScope struct{}"
	}

	return strings.Join([]string{"AuthObjectScope", string(data)}, " ")
}
