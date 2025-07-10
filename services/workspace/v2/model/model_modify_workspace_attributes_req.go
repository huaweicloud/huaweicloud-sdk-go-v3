package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModifyWorkspaceAttributesReq struct {
	AdInfo *AdDomainInfo `json:"ad_info,omitempty"`

	AdDomains *AdDomain `json:"ad_domains,omitempty"`

	// 接入模式。 - INTERNET：互联网接入。 - DEDICATED：专线接入。 - BOTH：代表两种接入方式都支持。
	AccessMode *ModifyWorkspaceAttributesReqAccessMode `json:"access_mode,omitempty"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	// ADN上网冲突网段列表，多个网段信息用分号隔开，列表长度不超过50。
	AdnConflictNetwork *string `json:"adn_conflict_network,omitempty"`

	// 子网的网络ID列表。
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 互联网接入端口。
	InternetAccessPort *string `json:"internet_access_port,omitempty"`

	// 企业ID。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 桌面退订是否发送邮件通知。
	IsSendEmail *bool `json:"is_send_email,omitempty"`

	// 开通专线访问VNC功能，如果传入的是default则自动创建，如果传入的自定义的dc_vnc_ip则直接使用，如果传入的是close表示关闭自定义VNC。
	DcVncIp *string `json:"dc_vnc_ip,omitempty"`

	// 是否授权收集日志。
	AuthorizedCollectLog *bool `json:"authorized_collect_log,omitempty"`

	// 是否授权hda升级。
	AuthorizedHdaUpgrade *bool `json:"authorized_hda_upgrade,omitempty"`

	ApplySharedVpcDedicatedParam *ApplySharedVpcDedicatedParam `json:"apply_shared_vpc_dedicated_param,omitempty"`

	ApplyDedicatedStandbyNetworkParam *ApplyDedicatedStandbyNetworkParam `json:"apply_dedicated_standby_network_param,omitempty"`

	// 是否授权桌面自动安装agent插件。
	IsAuthorizedInstallAgent *bool `json:"is_authorized_install_agent,omitempty"`

	// 是否授权最终租户创建快照。
	EnableUserCreateSnapshot *bool `json:"enable_user_create_snapshot,omitempty"`
}

func (o ModifyWorkspaceAttributesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyWorkspaceAttributesReq struct{}"
	}

	return strings.Join([]string{"ModifyWorkspaceAttributesReq", string(data)}, " ")
}

type ModifyWorkspaceAttributesReqAccessMode struct {
	value string
}

type ModifyWorkspaceAttributesReqAccessModeEnum struct {
	INTERNET  ModifyWorkspaceAttributesReqAccessMode
	DEDICATED ModifyWorkspaceAttributesReqAccessMode
	BOTH      ModifyWorkspaceAttributesReqAccessMode
}

func GetModifyWorkspaceAttributesReqAccessModeEnum() ModifyWorkspaceAttributesReqAccessModeEnum {
	return ModifyWorkspaceAttributesReqAccessModeEnum{
		INTERNET: ModifyWorkspaceAttributesReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: ModifyWorkspaceAttributesReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: ModifyWorkspaceAttributesReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c ModifyWorkspaceAttributesReqAccessMode) Value() string {
	return c.value
}

func (c ModifyWorkspaceAttributesReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyWorkspaceAttributesReqAccessMode) UnmarshalJSON(b []byte) error {
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
