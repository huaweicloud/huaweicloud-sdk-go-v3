package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTopAbnormalRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间（13位毫秒时间戳），需要和to同时使用

	From *int64 `json:"from,omitempty"`
	// 结束时间（13位毫秒时间戳），需要和from同时使用

	To *int64 `json:"to,omitempty"`
	// 要查询的前几的结果

	Top *int32 `json:"top,omitempty"`
	// 状态码

	Code *int32 `json:"code,omitempty"`
	// 要查询域名列表（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表（仅独享模式涉及）

	Instances *string `json:"instances,omitempty"`
}

func (o ListTopAbnormalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopAbnormalRequest struct{}"
	}

	return strings.Join([]string{"ListTopAbnormalRequest", string(data)}, " ")
}
