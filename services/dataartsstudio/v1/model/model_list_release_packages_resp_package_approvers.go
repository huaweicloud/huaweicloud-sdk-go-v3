package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListReleasePackagesRespPackageApprovers struct {

	// 发布包审批人id
	UserId *string `json:"user_id,omitempty"`

	// 发布包审批人名称
	UserName *string `json:"user_name,omitempty"`
}

func (o ListReleasePackagesRespPackageApprovers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReleasePackagesRespPackageApprovers struct{}"
	}

	return strings.Join([]string{"ListReleasePackagesRespPackageApprovers", string(data)}, " ")
}
