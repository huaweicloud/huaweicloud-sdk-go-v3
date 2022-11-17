package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListWorkspacesResponse struct {

	// 唯一标识ID。
	Id *string `json:"id,omitempty"`

	AdDomains *AdInfo `json:"ad_domains,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 接入方式。 - INTERNET：表示互联网接入。 - DEDICATED：表示专线接入。 - BOTH：表示同时支持互联网接入和专线接入。
	AccessMode *string `json:"access_mode,omitempty"`

	// 专线接入网段，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	// 专线接入地址，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedAccessAddress *string `json:"dedicated_access_address,omitempty"`

	// 互联网接入地址，只有access_mode为“INTERNET”或“BOTH”时才会返回该参数。
	InternetAccessAddress *string `json:"internet_access_address,omitempty"`

	// 互联网接入端口。
	InternetAccessPort *string `json:"internet_access_port,omitempty"`

	// 云办公服务的状态。 - PREPARING：准备开通。 - SUBSCRIBING：开通中。 - SUBSCRIBED：已开通。 - SUBSCRIPTION_FAILED：开通失败。 - DEREGISTERING：销户中。 - DEREGISTRATION_FAILED：销户失败。 - CLOSED：已销户未开通。
	Status *ListWorkspacesResponseStatus `json:"status,omitempty"`

	// 互联网和专线切换任务的状态。 - init： 初始化 - 开通服务后的初始状态。 - available： 可用 - 执行过任务且成功后恢复的正常状态。 - internetOpening： 开启中 - 开通互联网接入开启中。 - dedicatedOpening： 开启中 - 开通专线接入开启中。 - internetOpenFailed： 开启失败 - 开通互联网接入开启失败。 - dedicatedOpenFailed： 开启失败 - 开通专线接入开启失败。 - openSuccess： 开启成功 - 开通互联网接入成功。 - internetClosing： 关闭中 - 关闭互联网接入关闭中。 - dedicatedClosing： 关闭中 - 关闭专线接入关闭中。 - internetCloseFailed： 关闭失败 - 关闭互联网接入方式失败。 - dedicatedCloseFailed： 关闭失败 - 关闭专线接入方式失败。 - closeSuccess： 关闭成功 - 关闭接入方式成功。 - internetAccessPortModifying： 互联网接入端口修改中。 - internetAccessPortModifyFailed： 端口修改失败。
	AccessStatus *string `json:"access_status,omitempty"`

	// 业务子网，可以指定返回的网络ID订购桌面。
	SubnetIds *[]SubnetInfo `json:"subnet_ids,omitempty"`

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

	// 桌面退订是否发送邮件通知。
	IsSendEmail    *bool `json:"is_send_email,omitempty"`
	HttpStatusCode int   `json:"-"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
