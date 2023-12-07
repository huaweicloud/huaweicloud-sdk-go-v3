package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetFactoryJobTagsRequest Request Object
type SetFactoryJobTagsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	// 有Body体的情况下必选，无Body体的情况下则无需填写和校验
	ContentType *string `json:"Content-Type,omitempty"`

	// 使用AK/SK进行认证时该字段必选
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK进行认证时该字段必选
	Host *string `json:"Host,omitempty"`

	Body *SetJobTagsRequestBody `json:"body,omitempty"`
}

func (o SetFactoryJobTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetFactoryJobTagsRequest struct{}"
	}

	return strings.Join([]string{"SetFactoryJobTagsRequest", string(data)}, " ")
}
