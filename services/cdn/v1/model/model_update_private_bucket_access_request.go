package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePrivateBucketAccessRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属项目，不传表示查询默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 加速域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`

	Body *UpdatePrivateBucketAccessBody `json:"body,omitempty"`
}

func (o UpdatePrivateBucketAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateBucketAccessRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateBucketAccessRequest", string(data)}, " ")
}
