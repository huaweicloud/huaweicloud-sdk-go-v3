package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Host struct {

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: Agent ID **取值范围**: 字符长度1-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 私有IP地址 **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 弹性公网IP地址 **取值范围**: 字符长度1-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 企业项目ID **取值范围**: 字符长度0-256位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度0-256位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**: 系统名称 **取值范围**: 字符长度0-128位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 系统版本 **取值范围**: 字符长度0-256位
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释**: 内核版本 **取值范围**: 字符长度0-256位
	KernelVersion *string `json:"kernel_version,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: 包含如下4种。 - ACTIVE ：运行中。 - SHUTOFF ：关机。 - BUILDING ：创建中。 - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: Agent状态 **取值范围**: 包含如下6种。 - installed ：已安装。 - not_installed ：未安装。 - online ：在线。 - offline ：离线。 - install_failed ：安装失败。 - installing ：安装中。
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 安装结果 **取值范围**: 包含如下12种。   - install_succeed ：安装成功。   - network_access_timeout ：网络不通，访问超时。   - invalid_port ：无效端口。   - auth_failed ：认证错误，口令不正确。   - permission_denied ：权限错误，被拒绝。   - no_available_vpc ：没有相同VPC的agent在线虚拟机。   - install_exception ：安装异常。   - invalid_param ：参数错误。   - install_failed ：安装失败。   - package_unavailable ：安装包失效。   - os_type_not_support ：系统类型错误。   - os_arch_not_support ：架构类型错误。
	InstallResultCode *string `json:"install_result_code,omitempty"`

	// **参数解释**： 主机开通的版本 **取值范围**： 包含如下7种输入。 - hss.version.null ：无。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 防护状态 **取值范围**： 包含如下3种。 - closed ：未防护。 - opened ：防护中。 - protection_exception ：防护异常。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**： 系统镜像 **取值范围**： 字符长度0-128位
	OsImage *string `json:"os_image,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： 包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 操作系统位数 **取值范围**： 字符长度0-128位
	OsBit *string `json:"os_bit,omitempty"`

	// **参数解释**： 云主机安全检测结果 **取值范围**： 包含如下4种。 - undetected ：未检测。 - clean ：无风险。 - risk ：有风险。 - scanning ：检测中。
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释**： 试用版到期时间 **取值范围**： -1到4824695185000（-1表示非试用版配额，当值不为-1时为试用版本过期时间）
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释**： 收费模式 **取值范围**： 包含如下2种。   - packet_cycle ：包年/包月。   - on_demand ：按需。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**： 主机安全配额ID（UUID） **取值范围**： 字符长度0-128位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 是否非华为云机器 **取值范围**： true或者false
	OutsideHost *bool `json:"outside_host,omitempty"`

	// **参数解释**： 服务器组ID **取值范围**： 字符长度0-128位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 服务器组名称 **取值范围**： 字符长度0-128位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度0-128位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**： 策略组名称 **取值范围**： 字符长度0-128位
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释**： 资产风险 **取值范围**： 0-2097152
	Asset *int32 `json:"asset,omitempty"`

	// **参数解释**： 漏洞风险总数，包含Linux软件漏洞、Windows系统漏洞、Web-CMS漏洞、应用漏洞 **取值范围**： 0-2097152
	Vulnerability *int32 `json:"vulnerability,omitempty"`

	// **参数解释**： 基线风险总数，包含配置风险、弱口令 **取值范围**： 0-2097152
	Baseline *int32 `json:"baseline,omitempty"`

	// **参数解释**： 入侵风险总数 **取值范围**： 0-2097152
	Intrusion *int32 `json:"intrusion,omitempty"`

	// **参数解释**： 资产重要性 **取值范围**： 包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 标签列表 **取值范围**： 不涉及
	Labels *[]string `json:"labels,omitempty"`

	// **参数解释**： agent安装时间，采用时间戳，默认毫秒 **取值范围**： 0-4824695185000
	AgentCreateTime *int64 `json:"agent_create_time,omitempty"`

	// **参数解释**： agent状态修改时间，采用时间戳，默认毫秒 **取值范围**： 0-4824695185000
	AgentUpdateTime *int64 `json:"agent_update_time,omitempty"`

	// **参数解释**： agent版本 **取值范围**： 字符长度0-32位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**： 升级状态 **取值范围**： 包含如下4种。   - not_upgrade ：未升级，也就是默认状态，客户还没有给这台机器下发过升级。   - upgrading ：正在升级中。   - upgrade_failed ：升级失败。   - upgrade_succeed ：升级成功。
	UpgradeStatus *string `json:"upgrade_status,omitempty"`

	// **参数解释**： 升级失败原因，只有当 upgrade_status 为 upgrade_failed 时才显示 **取值范围**： 包含如下6种。   - package_unavailable ：升级包解析失败，升级文件有错误。   - network_access_timeout ：下载升级包失败，网络异常。   - agent_offline ：agent离线。   - hostguard_abnormal ：agent工作进程异常。   - insufficient_disk_space ：磁盘空间不足。   - failed_to_replace_file ：替换文件失败。
	UpgradeResultCode *string `json:"upgrade_result_code,omitempty"`

	// **参数解释**： 该服务器agent是否可升级 **取值范围**： true或者false
	Upgradable *bool `json:"upgradable,omitempty"`

	// **参数解释**： 开启防护时间，采用时间戳，默认毫秒 **取值范围**： 0-4824695185000
	OpenTime *int64 `json:"open_time,omitempty"`

	// **参数解释**： 防护是否中断 **取值范围**： true或者false
	ProtectInterrupt *bool `json:"protect_interrupt,omitempty"`

	// **参数解释**： 防护是否降级 **取值范围**： true或者false
	ProtectDegradation *bool `json:"protect_degradation,omitempty"`

	// **参数解释**： 服务器来源 **取值范围**： 包含如下3种。   - ecs ：华为云ecs。   - outside ：非华为云机器。   - workspace ：华为云workspace。
	HostSources *HostHostSources `json:"host_sources,omitempty"`

	// **参数解释**： 防护中断原因 **取值范围**： 字符长度1-32位
	InterruptReason *string `json:"interrupt_reason,omitempty"`

	// **参数解释**： 防护降级原因 **取值范围**： 字符长度1-32位
	DegradationReason *string `json:"degradation_reason,omitempty"`

	// **参数解释**： 使用的密钥对名称 **取值范围**： 字符长度1-32位
	KeyName *string `json:"key_name,omitempty"`

	// **参数解释**： cce购买主机 **取值范围**： 字符长度1-32位
	AutoOpenVersion *string `json:"auto_open_version,omitempty"`

	// **参数解释**： 安装进度 **取值范围**： 0-100
	InstallProgress *int32 `json:"install_progress,omitempty"`

	// **参数解释**： vpc id **取值范围**： 字符长度0-128位
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 后台识别服务器常用登录地编号 **取值范围**： 不涉及
	CommonLoginAreaCodes *[]int32 `json:"common_login_area_codes,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 字符长度1-128位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群id **取值范围**： 字符长度1-128位
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o Host) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Host struct{}"
	}

	return strings.Join([]string{"Host", string(data)}, " ")
}

type HostHostSources struct {
	value string
}

type HostHostSourcesEnum struct {
	ECS       HostHostSources
	OUTSIDE   HostHostSources
	WORKSPACE HostHostSources
}

func GetHostHostSourcesEnum() HostHostSourcesEnum {
	return HostHostSourcesEnum{
		ECS: HostHostSources{
			value: "ecs",
		},
		OUTSIDE: HostHostSources{
			value: "outside",
		},
		WORKSPACE: HostHostSources{
			value: "workspace",
		},
	}
}

func (c HostHostSources) Value() string {
	return c.value
}

func (c HostHostSources) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostHostSources) UnmarshalJSON(b []byte) error {
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
