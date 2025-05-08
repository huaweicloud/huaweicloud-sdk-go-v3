package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesBatchRequest Request Object
type ListInstancesBatchRequest struct {

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文。
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID。
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息。
	XUserProfile *string `json:"x-user-profile,omitempty"`

	Body *InstancesBatchesMode `json:"body,omitempty"`
}

func (o ListInstancesBatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesBatchRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesBatchRequest", string(data)}, " ")
}
