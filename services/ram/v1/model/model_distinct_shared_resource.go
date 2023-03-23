package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 描述共享的不同资源。
type DistinctSharedResource struct {

	// 资源的统一资源名称。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 资源的类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 最后一次更新资源共享实例的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o DistinctSharedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinctSharedResource struct{}"
	}

	return strings.Join([]string{"DistinctSharedResource", string(data)}, " ")
}
