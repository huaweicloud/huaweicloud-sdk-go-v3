package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFactoryJobNameRequest Request Object
type UpdateFactoryJobNameRequest struct {

	// 作业名
	JobName string `json:"job_name"`

	// 工作空间ID
	Workspace string `json:"workspace"`

	// 有Body体的情况下必选，无Body体的情况下则无需填写和校验
	ContentType *string `json:"Content-Type,omitempty"`

	// 使用AK/SK进行认证时该字段必选
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK进行认证时该字段必选
	Host *string `json:"Host,omitempty"`

	Body *UpdateFactoryJobNameRequestBody `json:"body,omitempty"`
}

func (o UpdateFactoryJobNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFactoryJobNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateFactoryJobNameRequest", string(data)}, " ")
}
