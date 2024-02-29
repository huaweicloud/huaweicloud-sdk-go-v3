package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryReleasePackagesRequest Request Object
type ListFactoryReleasePackagesRequest struct {

	// 工作空间ID，默认查询default空间
	Workspace *string `json:"workspace,omitempty"`

	// 有Body体的情况下必须，无Body体的情况下则无需填写和校验，默认值：application/json
	ContentType *string `json:"Content-Type,omitempty"`

	Body *ListReleasePackagesRequestBody `json:"body,omitempty"`
}

func (o ListFactoryReleasePackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryReleasePackagesRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryReleasePackagesRequest", string(data)}, " ")
}
