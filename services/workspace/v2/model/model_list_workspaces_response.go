package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkspacesResponse Response Object
type ListWorkspacesResponse struct {

	// 唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 主认证方式。 - KERBEROS：KERBEROS。 - KERBEROS_THIRD_SSO：第三方登录认证。
	AuthType *string `json:"auth_type,omitempty"`

	AdDomains *AdInfo `json:"ad_domains,omitempty"`

	ThirdGatewayInfo *ThirdGatewayInfo `json:"third_gateway_info,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 接入方式。 - INTERNET：表示互联网接入。 - DEDICATED：表示专线接入。 - BOTH：表示同时支持互联网接入和专线接入。
	AccessMode *string `json:"access_mode,omitempty"`

	// ADN上网冲突网段列表，多个网段信息用分号隔开，列表长度不超过50。
	AdnConflictNetwork *string `json:"adn_conflict_network,omitempty"`

	// 专线接入网段，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	// 专线接入地址，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedAccessAddress *string `json:"dedicated_access_address,omitempty"`

	// 专线接入ipv6地址，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedAccessAddressIpv6 *string `json:"dedicated_access_address_ipv6,omitempty"`

	// 互联网接入地址，只有access_mode为“INTERNET”或“BOTH”时才会返回该参数。
	InternetAccessAddress *string `json:"internet_access_address,omitempty"`

	// 互联网接入ipv6地址，只有access_mode为“INTERNET”或“BOTH”时才会返回该参数。
	InternetAccessAddressIpv6 *string `json:"internet_access_address_ipv6,omitempty"`

	// 互联网接入端口。
	InternetAccessPort *string `json:"internet_access_port,omitempty"`

	// 云办公服务的状态。 - PREPARING：准备开通。 - SUBSCRIBING：开通中。 - SUBSCRIBED：已开通。 - SUBSCRIPTION_FAILED：开通失败。 - DEREGISTERING：销户中。 - DEREGISTRATION_FAILED：销户失败。 - RECYCLING：清理资源中。 - RECYCLED：清理资源成功。 - RECYCLE_FAILED：清理资源失败。 - CLOSED：已销户未开通。
	Status *ListWorkspacesResponseStatus `json:"status,omitempty"`

	// 互联网和专线切换任务的状态。 - init： 初始化 - 开通服务后的初始状态。 - available： 可用 - 执行过任务且成功后恢复的正常状态。 - internetOpening： 开启中 - 开通互联网接入开启中。 - dedicatedOpening： 开启中 - 开通专线接入开启中。 - internetOpenFailed： 开启失败 - 开通互联网接入开启失败。 - dedicatedOpenFailed： 开启失败 - 开通专线接入开启失败。 - openSuccess： 开启成功 - 开通接入方式成功。 - internetClosing： 关闭中 - 关闭互联网接入关闭中。 - dedicatedClosing： 关闭中 - 关闭专线接入关闭中。 - internetCloseFailed： 关闭失败 - 关闭互联网接入方式失败。 - dedicatedCloseFailed： 关闭失败 - 关闭专线接入方式失败。 - closeSuccess： 关闭成功 - 关闭接入方式成功。 - internetAccessPortModifying： 互联网接入端口修改中。 - internetAccessPortModifyFailed： 端口修改失败。
	AccessStatus *string `json:"access_status,omitempty"`

	// 业务子网，可以指定返回的网络ID订购桌面。
	SubnetIds *[]SubnetInfo `json:"subnet_ids,omitempty"`

	// VPC配置信息列表。
	VpcConfigInfos *[]VpcConfigInfo `json:"vpc_config_infos,omitempty"`

	// 管理组件的子网网段。
	ManagementSubnetCidr *string `json:"management_subnet_cidr,omitempty"`

	InfrastructureSecurityGroup *SecurityGroup `json:"infrastructure_security_group,omitempty"`

	DesktopSecurityGroup *SecurityGroup `json:"desktop_security_group,omitempty"`

	// 是否可以取消服务。
	Closable *bool `json:"closable,omitempty"`

	// 配置状态。 - \"0\"： 开通服务成功，且对接AD成功。 - \"1\"： 开通服务成功，但AD配置失败。 - \"2\"： 开通服务成功，但AD配置失败后存在其他错误。 - \"3\"： 开通服务成功，AD未开启对接。
	ConfigStatus *string `json:"config_status,omitempty"`

	// 开通服务或注销服务的进度，格式为百分比，如：100%。
	Progress *string `json:"progress,omitempty"`

	// 开通服务或取消服务的任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 失败错误码。
	FailCode *int32 `json:"fail_code,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 企业ID。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源池类型。 - \"public\"： 私有资源池 - \"private\"： 公有资源池
	ProjectResourceType *string `json:"project_resource_type,omitempty"`

	AssistAuthConfigInfo *AssistAuthConfigInfo `json:"assist_auth_config_info,omitempty"`

	// 桌面退订是否发送邮件通知。
	IsSendEmail *bool `json:"is_send_email,omitempty"`

	// 是否授权收集日志。
	AuthorizedCollectLog *bool `json:"authorized_collect_log,omitempty"`

	// 是否授权hda升级。
	AuthorizedHdaUpgrade *bool `json:"authorized_hda_upgrade,omitempty"`

	// 站点配置
	SiteConfigs *[]SiteConfigsResponse `json:"site_configs,omitempty"`

	// 是否支持多VPC。
	IsMultiVpc *bool `json:"is_multi_vpc,omitempty"`

	// 是否支持配置nat映射。
	IsConfigNatMapping *bool `json:"is_config_nat_mapping,omitempty"`

	// 自定义的专线VNC地址
	DcVncIp *string `json:"dc_vnc_ip,omitempty"`

	// 专线VNC VPC终端节点ID
	DcVncVpcepId *string `json:"dc_vnc_vpcep_id,omitempty"`

	// 是否授权桌面自动安装agent插件。
	IsAuthorizedInstallAgent *bool `json:"is_authorized_install_agent,omitempty"`

	// 是否支持ipv6。
	IsSupportIpv6 *bool `json:"is_support_ipv6,omitempty"`

	// 是否授权最终租户创建快照。
	EnableUserCreateSnapshot *bool `json:"enable_user_create_snapshot,omitempty"`
	HttpStatusCode           int   `json:"-"`
}

func (o ListWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResponse", string(data)}, " ")
}

type ListWorkspacesResponseStatus struct {
	value string
}

type ListWorkspacesResponseStatusEnum struct {
	PREPARING             ListWorkspacesResponseStatus
	SUBSCRIBING           ListWorkspacesResponseStatus
	SUBSCRIBED            ListWorkspacesResponseStatus
	SUBSCRIPTION_FAILED   ListWorkspacesResponseStatus
	DEREGISTERING         ListWorkspacesResponseStatus
	DEREGISTRATION_FAILED ListWorkspacesResponseStatus
	RECYCLING             ListWorkspacesResponseStatus
	RECYCLED              ListWorkspacesResponseStatus
	RECYCLE_FAILED        ListWorkspacesResponseStatus
	CLOSED                ListWorkspacesResponseStatus
}

func GetListWorkspacesResponseStatusEnum() ListWorkspacesResponseStatusEnum {
	return ListWorkspacesResponseStatusEnum{
		PREPARING: ListWorkspacesResponseStatus{
			value: "PREPARING",
		},
		SUBSCRIBING: ListWorkspacesResponseStatus{
			value: "SUBSCRIBING",
		},
		SUBSCRIBED: ListWorkspacesResponseStatus{
			value: "SUBSCRIBED",
		},
		SUBSCRIPTION_FAILED: ListWorkspacesResponseStatus{
			value: "SUBSCRIPTION_FAILED",
		},
		DEREGISTERING: ListWorkspacesResponseStatus{
			value: "DEREGISTERING",
		},
		DEREGISTRATION_FAILED: ListWorkspacesResponseStatus{
			value: "DEREGISTRATION_FAILED",
		},
		RECYCLING: ListWorkspacesResponseStatus{
			value: "RECYCLING",
		},
		RECYCLED: ListWorkspacesResponseStatus{
			value: "RECYCLED",
		},
		RECYCLE_FAILED: ListWorkspacesResponseStatus{
			value: "RECYCLE_FAILED",
		},
		CLOSED: ListWorkspacesResponseStatus{
			value: "CLOSED",
		},
	}
}

func (c ListWorkspacesResponseStatus) Value() string {
	return c.value
}

func (c ListWorkspacesResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkspacesResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
