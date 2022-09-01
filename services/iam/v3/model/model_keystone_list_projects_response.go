package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListProjectsResponse struct {
	Links *Links `json:"links,omitempty" xml:"links"`

	// 项目信息列表。
	Projects       *[]ProjectResult `json:"projects,omitempty" xml:"projects"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListProjectsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectsResponse", string(data)}, " ")
}
