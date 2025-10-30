package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionServerInfo struct {

	// **参数解释**: 项目ID **取值范围**: 字符长度0-128
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 企业项目ID **取值范围**: 字符长度0-128
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器ID **取值范围**: 字符长度0-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: Agent ID **取值范围**: 字符长度0-128
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度0-128
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 弹性公网IP地址 **取值范围**: 字符长度0-128
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 私有IP地址 **取值范围**: 字符长度0-128
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 操作系统类型 **取值范围**:   包含如下2种。     - Linux ：Linux。     - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 系统名称 **取值范围**: 字符长度0-128
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: 包含如下2种。   - ACTIVE ：运行中。   - SHUTOFF ：关机。
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: 勒索防护状态 **取值范围**: 包含如下6种。   - closed ：未开启。   - opened ：防护中。   - opening ：开启中。   - closing ：关闭中。   - protect_failed：防护失败。   - protect_degraded：防护降级
	RansomProtectionStatus *string `json:"ransom_protection_status,omitempty"`

	// **参数解释**: 勒索防护失败细分原因 **取值范围**: 包含如下4种。   - driver_load_failed ：驱动加载失败。   - protect_interrupted ：防护中断。   - decoy_deploy_totally_failed ：全部诱饵部署失败。   - decoy_deploy_partially_failed ：部分诱饵部署失败。
	RansomProtectionFailReason *string `json:"ransom_protection_fail_reason,omitempty"`

	// **参数解释**: 诱饵防护失败的目录（仅部分诱饵部署失败状态有值） **取值范围**: 字符长度0-512
	FailedDecoyDir *string `json:"failed_decoy_dir,omitempty"`

	// **参数解释**: agent版本 **取值范围**: 字符长度1-128
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: 包含如下2种。 - closed ：未防护。 - opened ：防护中。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 服务器组ID **取值范围**: 字符长度1-128
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度1-128
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 防护策略ID **取值范围**: 字符长度1-128
	ProtectPolicyId *string `json:"protect_policy_id,omitempty"`

	// **参数解释**: 防护策略名称 **取值范围**: 字符长度1-128
	ProtectPolicyName *string `json:"protect_policy_name,omitempty"`

	BackupError *ProtectionServerInfoBackupError `json:"backup_error,omitempty"`

	// **参数解释**: 是否开启备份 **取值范围**: 包含如下3种。   - failed_to_turn_on_backup: 无法开启备份   - closed ：关闭。   - opened ：开启。
	BackupProtectionStatus *string `json:"backup_protection_status,omitempty"`

	// **参数解释**: 防护事件数 **取值范围**: 取值0-2097152
	CountProtectEvent *int32 `json:"count_protect_event,omitempty"`

	// **参数解释**: 已有备份数 **取值范围**: 取值0-2097152
	CountBackuped *int32 `json:"count_backuped,omitempty"`

	// **参数解释**: Agent状态 **取值范围**:   - installed：已安装。已安装状态包含以下四种情况：   - online：在线。表示Agent已经成功安装并且与HSS云端防护中心保持连接。   - offline：离线。表示虽然Agent已经安装，但当前与HSS云端防护中心的连接中断。   - install_failed：安装失败。表示在尝试安装过程中遇到错误或问题，导致安装未能完成。   - installing：安装中。表示当前正在进行Agent安装。   - not_installed ：未安装。表示服务器中尚未安装Agent。
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 主机开通的版本    **取值范围**: 包含如下7种输入：   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.advanced ：专业版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。   - hss.version.container.enterprise ：容器版。
	Version *string `json:"version,omitempty"`

	// **参数解释**: 服务器类型 **取值范围**: 包含如下3种输入：   - ecs ：弹性云服务器。   - outside ：线下主机。   - workspace ：云桌面。
	HostSource *string `json:"host_source,omitempty"`

	// **参数解释**: 存储库ID **取值范围**: 字符长度0-128
	VaultId *string `json:"vault_id,omitempty"`

	// **参数解释**: 存储库名称 **取值范围**: 字符长度0-128
	VaultName *string `json:"vault_name,omitempty"`

	// **参数解释**: 总容量，单位GB **取值范围**: 取值0-2097152
	VaultSize *int32 `json:"vault_size,omitempty"`

	// **参数解释**: 已使用容量，单位MB **取值范围**: 取值0-2097152
	VaultUsed *int32 `json:"vault_used,omitempty"`

	// **参数解释**: 已分配容量，单位GB，指绑定的服务器大小 **取值范围**: 取值0-2097152
	VaultAllocated *int32 `json:"vault_allocated,omitempty"`

	// **参数解释**: 存储库创建模式 **取值范围**: 包含如下2种： - 按需：post_paid - 包周期：pre_paid
	VaultChargingMode *string `json:"vault_charging_mode,omitempty"`

	// **参数解释**: 存储库状态。 **取值范围**: 包含如下5种：   - available ：可用。   - lock ：被锁定。   - frozen：冻结。   - deleting：删除中。   - error：错误。
	VaultStatus *string `json:"vault_status,omitempty"`

	// **参数解释**: 备份策略ID，若为空，则为未绑定状态，若不为空，通过backup_policy_enabled字段判断策略是否启用。 **取值范围**: 字符长度1-128
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	// **参数解释**: 备份策略名称 **取值范围**: 字符长度1-128
	BackupPolicyName *string `json:"backup_policy_name,omitempty"`

	// **参数解释**: 策略是否启用 **取值范围**: 包含如下2种：   - true ：策略已启用。   - false ：策略未启用。
	BackupPolicyEnabled *bool `json:"backup_policy_enabled,omitempty"`

	// **参数解释**: 已绑定服务器（个） **取值范围**: 取值0-2097152
	ResourcesNum *int32 `json:"resources_num,omitempty"`
}

func (o ProtectionServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionServerInfo struct{}"
	}

	return strings.Join([]string{"ProtectionServerInfo", string(data)}, " ")
}
