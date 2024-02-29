package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryPackageDetailResponse Response Object
type ShowFactoryPackageDetailResponse struct {
	ReleasePackage *ShowPackageDetailRespReleasePackage `json:"release_package,omitempty"`

	// 发布任务详情
	TaskDetails *[]ShowPackageDetailRespTaskDetails `json:"task_details,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFactoryPackageDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryPackageDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFactoryPackageDetailResponse", string(data)}, " ")
}
