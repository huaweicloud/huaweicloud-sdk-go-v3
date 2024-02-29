package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPackageDetailRespReleasePackagePackageApprovers struct {

	// 发布包审批人id
	UserId *string `json:"user_id,omitempty"`

	// 发布包审批人id
	UserName *string `json:"user_name,omitempty"`
}

func (o ShowPackageDetailRespReleasePackagePackageApprovers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageDetailRespReleasePackagePackageApprovers struct{}"
	}

	return strings.Join([]string{"ShowPackageDetailRespReleasePackagePackageApprovers", string(data)}, " ")
}
