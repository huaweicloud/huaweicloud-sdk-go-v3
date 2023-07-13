package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsAppQueryResp struct {

	// 应用主键ID，用于获取、修改应用的唯一标识
	Id *string `json:"id,omitempty"`

	// 创建时间[yyyy-MM-dd HH:mm:ss]
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间[yyyy-MM-dd HH:mm:ss]
	UpdateTime *string `json:"update_time,omitempty"`

	// 租户customer id
	CustomerId *string `json:"customer_id,omitempty"`

	// 租户resource id
	ResourceId *string `json:"resource_id,omitempty"`

	// 租户开发者账号
	DeveloperAccount *string `json:"developer_account,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// omp应用名称
	OmpAppName *string `json:"omp_app_name,omitempty"`

	// 应用key
	AppKey *string `json:"app_key,omitempty"`

	// 上行短信地址
	UpLinkAddr *string `json:"up_link_addr,omitempty"`

	// 应用状态   CREATED：待上线。应用暂未创建成功，请稍候。   SUSPENDED：暂停。无法发起业务请求。当客户所发短信内容触发业务违规，或客户申请退订短信业务时，运营经理会将客户短信应用暂停。   LAUNCHED：正常。应用添加成功，可以正常使用。
	Status *string `json:"status,omitempty"`

	// 行业类型
	Industry *int32 `json:"industry,omitempty"`

	// 地域 1. cn：国内 2. intl：国际
	Region *string `json:"region,omitempty"`

	// 国际/港澳台短信通道号
	IntlChannelNum *string `json:"intl_channel_num,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// IP白名单
	IpWhiteList *string `json:"ip_white_list,omitempty"`

	// 接入地址
	AppAccessAddr *string `json:"app_access_addr,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 平台
	Platform *string `json:"platform,omitempty"`

	// 是否支持多OMP
	IsSupportMultiomp *bool `json:"is_support_multiomp,omitempty"`

	Tenant *TenantBasic `json:"tenant,omitempty"`
}

func (o SmsAppQueryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsAppQueryResp struct{}"
	}

	return strings.Join([]string{"SmsAppQueryResp", string(data)}, " ")
}
