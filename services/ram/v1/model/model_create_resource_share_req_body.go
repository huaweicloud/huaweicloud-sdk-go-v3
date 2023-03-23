package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// The request body of the CreateResourceShare operation.
type CreateResourceShareReqBody struct {

	// 资源共享实例的名称。
	Name string `json:"name"`

	// 资源共享实例的描述。
	Description *string `json:"description,omitempty"`

	// 资源共享实例关联的RAM权限列表。一种资源类型只能关联一个RAM权限。如果您没有指定权限ID，RAM将自动为每个资源类型关联默认权限。
	PermissionIds *[]string `json:"permission_ids,omitempty"`

	// 资源共享实例关联的一个或多个资源使用者的列表。
	Principals *[]string `json:"principals,omitempty"`

	// 资源共享实例关联的一个或多个共享资源URN的列表。
	ResourceUrns *[]string `json:"resource_urns,omitempty"`

	// 资源共享标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateResourceShareReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceShareReqBody struct{}"
	}

	return strings.Join([]string{"CreateResourceShareReqBody", string(data)}, " ")
}
