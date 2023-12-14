package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudTableClusterRequest Request Object
type CreateCloudTableClusterRequest struct {

	// 语言类型
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateClusterRequestBody `json:"body,omitempty"`
}

func (o CreateCloudTableClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudTableClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudTableClusterRequest", string(data)}, " ")
}
