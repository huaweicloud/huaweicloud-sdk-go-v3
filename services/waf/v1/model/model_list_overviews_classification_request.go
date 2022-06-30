package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListOverviewsClassificationRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 起始时间（13位毫秒时间戳），需要和to同时使用
	From int64 `json:"from"`

	// 结束时间（13位毫秒时间戳），需要和from同时使用
	To int64 `json:"to"`

	// 要查询的前几的结果，最大值为10
	Top *int32 `json:"top,omitempty"`

	// 要查询域名列表
	Hosts *string `json:"hosts,omitempty"`

	// 要查询实例列表
	Instances *string `json:"instances,omitempty"`
}

func (o ListOverviewsClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOverviewsClassificationRequest struct{}"
	}

	return strings.Join([]string{"ListOverviewsClassificationRequest", string(data)}, " ")
}
