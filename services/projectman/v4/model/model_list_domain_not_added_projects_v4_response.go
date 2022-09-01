package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainNotAddedProjectsV4Response struct {

	// 项目信息列表
	Projects *[]ListDomainNotAddedProjectsV4ResponseBodyProjects `json:"projects,omitempty" xml:"projects"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDomainNotAddedProjectsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainNotAddedProjectsV4Response struct{}"
	}

	return strings.Join([]string{"ListDomainNotAddedProjectsV4Response", string(data)}, " ")
}
