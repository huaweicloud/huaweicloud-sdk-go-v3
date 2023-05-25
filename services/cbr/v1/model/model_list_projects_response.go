package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectsResponse struct {

	// 项目信息
	Projects *[]ProjectsListInfo `json:"projects,omitempty"`

	Links          *SelfLinksInfo `json:"links,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectsResponse", string(data)}, " ")
}
