package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListAuthProjectsResponse struct {
	Links *LinksSelf `json:"links,omitempty" xml:"links"`

	// 项目信息列表。
	Projects       *[]AuthProjectResult `json:"projects,omitempty" xml:"projects"`
	HttpStatusCode int                  `json:"-"`
}

func (o KeystoneListAuthProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListAuthProjectsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthProjectsResponse", string(data)}, " ")
}
