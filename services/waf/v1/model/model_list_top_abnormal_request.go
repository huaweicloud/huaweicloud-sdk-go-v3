package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopAbnormalRequest Request Object
type ListTopAbnormalRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 起始时间（13位毫秒时间戳），需要和to同时使用
	From int64 `json:"from"`

	// 结束时间（13位毫秒时间戳），需要和from同时使用
	To int64 `json:"to"`

	// 要查询的前几的结果，默认值为5，最大值为10。
	Top *int32 `json:"top,omitempty"`

	// 要查询的异常状态码，目前支持查询的异常状态码包括404、500以及502。不传该参数默认查询404的状态码。
	Code *int32 `json:"code,omitempty"`

	// 域名id，通过查询云模式防护域名列表（ListHost）获取域名id或者通过独享模式域名列表（ListPremiumHost）获取域名id。默认不传，查询该项目下所有防护域名的top业务异常统计信息。
	Hosts *string `json:"hosts,omitempty"`

	// 要查询引擎实例id
	Instances *string `json:"instances,omitempty"`
}

func (o ListTopAbnormalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopAbnormalRequest struct{}"
	}

	return strings.Join([]string{"ListTopAbnormalRequest", string(data)}, " ")
}
