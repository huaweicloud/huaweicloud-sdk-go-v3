package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DistinctSharedPrincipal 描述共享的不同角色。
type DistinctSharedPrincipal struct {

	// 资源共享实例的创建者或使用者的帐号ID或URN。
	Id *string `json:"id,omitempty"`

	// 最后一次更新资源共享实例的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o DistinctSharedPrincipal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinctSharedPrincipal struct{}"
	}

	return strings.Join([]string{"DistinctSharedPrincipal", string(data)}, " ")
}
