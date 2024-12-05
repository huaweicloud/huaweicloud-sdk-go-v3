package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListFederationProjectsResponse Response Object
type KeystoneListFederationProjectsResponse struct {
	Links *LinksSelf `json:"links,omitempty"`

	// 项目信息列表。
	Projects       *[]AuthProjectResult `json:"projects,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o KeystoneListFederationProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListFederationProjectsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListFederationProjectsResponse", string(data)}, " ")
}
