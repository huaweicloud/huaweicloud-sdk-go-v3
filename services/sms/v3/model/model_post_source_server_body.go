package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostSourceServerBody 源端服务器
type PostSourceServerBody struct {

	// 源端在SMS数据库中的ID
	Id *string `json:"id,omitempty"`

	// 源端服务器ip，格式需满足ip标准格式。ip与ipv6必填一个。
	Ip *string `json:"ip,omitempty"`

	// 源端服务器ip，格式需满足ipv6标准格式。ip与ipv6必填一个。
	Ipv6 *string `json:"ipv6,omitempty"`

	// 用来区分不同源端服务器的名称
	Name string `json:"name"`

	// 源端主机名
	Hostname *string `json:"hostname,omitempty"`

	// 源端服务器的OS类型，分为Windows和Linux，注册必选，更新非必选
	OsType PostSourceServerBodyOsType `json:"os_type"`

	// 操作系统版本，注册必选，更新非必选
	OsVersion string `json:"os_version"`

	// Linux操作系统块检查
	LinuxBlockCheck *string `json:"linux_block_check,omitempty"`

	// 源端服务器启动类型，如BIOS或者UEFI
	Firmware *PostSourceServerBodyFirmware `json:"firmware,omitempty"`

	// CPU个数，单位vCPU
	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	// 内存大小，单位MB
	Memory *int64 `json:"memory,omitempty"`

	// 源端服务器的磁盘信息
	Disks *[]ServerDisk `json:"disks,omitempty"`

	// 源端的Btrfs信息。如果源端不存在Btrfs，则为[] Linux场景必选，否则无法通过后续环境检查
	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`

	// 源端服务器的网卡信息
	Networks []NetWork `json:"networks"`

	// 租户的domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 是否安装rsync组件，Linux系统此参数为必选，否则无法通过后续环境检查
	HasRsync *bool `json:"has_rsync,omitempty"`

	// 源端是否是半虚拟化 Linux场景必选，否则无法通过后续环境检查
	Paravirtualization *bool `json:"paravirtualization,omitempty"`

	// 裸设备列表 Linux场景必选，否则无法通过后续环境检查
	RawDevices *string `json:"raw_devices,omitempty"`

	// 是否缺少驱动文件 Windows场景必选，否则无法通过后续环境检查
	DriverFiles *bool `json:"driver_files,omitempty"`

	// 是否存在不正常服务 Windows场景必选，否则无法通过后续环境检查
	SystemServices *bool `json:"system_services,omitempty"`

	// 权限是否满足要求 Windows场景必选，否则无法通过后续环境检查
	AccountRights *bool `json:"account_rights,omitempty"`

	// 系统引导类型 仅允许“GRUB”取值，Linux场景必选，否则无法通过后续环境检查
	BootLoader *string `json:"boot_loader,omitempty"`

	// 系统目录 Windows场景必选，否则无法通过后续环境检查
	SystemDir *string `json:"system_dir,omitempty"`

	// 卷组 如果没有卷组，输入[] Linux场景必选，否则无法通过后续环境检查
	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`

	// Agent版本
	AgentVersion string `json:"agent_version"`

	// 内核版本信息
	KernelVersion *string `json:"kernel_version,omitempty"`

	// 迁移周期 cutovering:启动目的端中 cutovered:启动目的端完成 checking:检查中 setting:设置中 replicating:复制中 syncing:同步中
	MigrationCycle *PostSourceServerBodyMigrationCycle `json:"migration_cycle,omitempty"`

	// 源端服务器状态 unavailable：环境校验不通过 waiting：等待 initialize：初始化 replicate：复制 syncing：持续同步 stopping：暂停中 stopped：已暂停 skipping：跳过中 deleting：删除中 clearing: 清理快照资源中 cleared：清理快照资源完成 clearfailed：清理快照资源失败 premigready：迁移演练就绪 premiging：迁移演练中 premiged：迁移演练完成 premigfailed：迁移演练失败 cloning：等待克隆完成 cutovering：启动目的端中 finished：启动目的端完成 error：错误
	State *PostSourceServerBodyState `json:"state,omitempty"`

	// 是否是OEM操作系统(Windows)
	OemSystem *bool `json:"oem_system,omitempty"`

	// 启动方式 可以取值MANUAL、MGC或者空，不进行校验
	StartType *string `json:"start_type,omitempty"`

	// 磁盘IO读时延，单位为ms
	IoReadWait *float64 `json:"io_read_wait,omitempty"`

	// 是否安装tc组件，Linux系统此参数为必选，否则无法通过后续环境检查
	HasTc *bool `json:"has_tc,omitempty"`

	// 平台信息: hw：华为  ali：阿里 aws：亚马逊 azure：微软云 gcp：谷歌云 tencent：腾讯云 vmware：VMware hyperv：HyperV other：其他 default：默认
	Platform *PostSourceServerBodyPlatform `json:"platform,omitempty"`
}

func (o PostSourceServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostSourceServerBody struct{}"
	}

	return strings.Join([]string{"PostSourceServerBody", string(data)}, " ")
}

type PostSourceServerBodyOsType struct {
	value string
}

type PostSourceServerBodyOsTypeEnum struct {
	WINDOWS PostSourceServerBodyOsType
	LINUX   PostSourceServerBodyOsType
}

func GetPostSourceServerBodyOsTypeEnum() PostSourceServerBodyOsTypeEnum {
	return PostSourceServerBodyOsTypeEnum{
		WINDOWS: PostSourceServerBodyOsType{
			value: "WINDOWS",
		},
		LINUX: PostSourceServerBodyOsType{
			value: "LINUX",
		},
	}
}

func (c PostSourceServerBodyOsType) Value() string {
	return c.value
}

func (c PostSourceServerBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyOsType) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyFirmware struct {
	value string
}

type PostSourceServerBodyFirmwareEnum struct {
	BIOS PostSourceServerBodyFirmware
	UEFI PostSourceServerBodyFirmware
}

func GetPostSourceServerBodyFirmwareEnum() PostSourceServerBodyFirmwareEnum {
	return PostSourceServerBodyFirmwareEnum{
		BIOS: PostSourceServerBodyFirmware{
			value: "BIOS",
		},
		UEFI: PostSourceServerBodyFirmware{
			value: "UEFI",
		},
	}
}

func (c PostSourceServerBodyFirmware) Value() string {
	return c.value
}

func (c PostSourceServerBodyFirmware) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyFirmware) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyMigrationCycle struct {
	value string
}

type PostSourceServerBodyMigrationCycleEnum struct {
	CUTOVERING  PostSourceServerBodyMigrationCycle
	CUTOVERED   PostSourceServerBodyMigrationCycle
	CHECKING    PostSourceServerBodyMigrationCycle
	SETTING     PostSourceServerBodyMigrationCycle
	REPLICATING PostSourceServerBodyMigrationCycle
	SYNCING     PostSourceServerBodyMigrationCycle
}

func GetPostSourceServerBodyMigrationCycleEnum() PostSourceServerBodyMigrationCycleEnum {
	return PostSourceServerBodyMigrationCycleEnum{
		CUTOVERING: PostSourceServerBodyMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: PostSourceServerBodyMigrationCycle{
			value: "cutovered",
		},
		CHECKING: PostSourceServerBodyMigrationCycle{
			value: "checking",
		},
		SETTING: PostSourceServerBodyMigrationCycle{
			value: "setting",
		},
		REPLICATING: PostSourceServerBodyMigrationCycle{
			value: "replicating",
		},
		SYNCING: PostSourceServerBodyMigrationCycle{
			value: "syncing",
		},
	}
}

func (c PostSourceServerBodyMigrationCycle) Value() string {
	return c.value
}

func (c PostSourceServerBodyMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyMigrationCycle) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyState struct {
	value string
}

type PostSourceServerBodyStateEnum struct {
	UNAVAILABLE  PostSourceServerBodyState
	WAITING      PostSourceServerBodyState
	INITIALIZE   PostSourceServerBodyState
	REPLICATE    PostSourceServerBodyState
	SYNCING      PostSourceServerBodyState
	STOPPING     PostSourceServerBodyState
	STOPPED      PostSourceServerBodyState
	SKIPPING     PostSourceServerBodyState
	DELETING     PostSourceServerBodyState
	CLEARING     PostSourceServerBodyState
	CLEARED      PostSourceServerBodyState
	CLEARFAILED  PostSourceServerBodyState
	PREMIGREADY  PostSourceServerBodyState
	PREMIGED     PostSourceServerBodyState
	PREMIGFAILED PostSourceServerBodyState
	CLONING      PostSourceServerBodyState
	CUTOVERING   PostSourceServerBodyState
	FINISHED     PostSourceServerBodyState
	ERROR        PostSourceServerBodyState
}

func GetPostSourceServerBodyStateEnum() PostSourceServerBodyStateEnum {
	return PostSourceServerBodyStateEnum{
		UNAVAILABLE: PostSourceServerBodyState{
			value: "unavailable",
		},
		WAITING: PostSourceServerBodyState{
			value: "waiting",
		},
		INITIALIZE: PostSourceServerBodyState{
			value: "initialize",
		},
		REPLICATE: PostSourceServerBodyState{
			value: "replicate",
		},
		SYNCING: PostSourceServerBodyState{
			value: "syncing",
		},
		STOPPING: PostSourceServerBodyState{
			value: "stopping",
		},
		STOPPED: PostSourceServerBodyState{
			value: "stopped",
		},
		SKIPPING: PostSourceServerBodyState{
			value: "skipping",
		},
		DELETING: PostSourceServerBodyState{
			value: "deleting",
		},
		CLEARING: PostSourceServerBodyState{
			value: "clearing",
		},
		CLEARED: PostSourceServerBodyState{
			value: "cleared",
		},
		CLEARFAILED: PostSourceServerBodyState{
			value: "clearfailed",
		},
		PREMIGREADY: PostSourceServerBodyState{
			value: "premigready",
		},
		PREMIGED: PostSourceServerBodyState{
			value: "premiged",
		},
		PREMIGFAILED: PostSourceServerBodyState{
			value: "premigfailed",
		},
		CLONING: PostSourceServerBodyState{
			value: "cloning",
		},
		CUTOVERING: PostSourceServerBodyState{
			value: "cutovering",
		},
		FINISHED: PostSourceServerBodyState{
			value: "finished",
		},
		ERROR: PostSourceServerBodyState{
			value: "error",
		},
	}
}

func (c PostSourceServerBodyState) Value() string {
	return c.value
}

func (c PostSourceServerBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyState) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyPlatform struct {
	value string
}

type PostSourceServerBodyPlatformEnum struct {
	HW      PostSourceServerBodyPlatform
	ALI     PostSourceServerBodyPlatform
	AWS     PostSourceServerBodyPlatform
	AZURE   PostSourceServerBodyPlatform
	GCP     PostSourceServerBodyPlatform
	TENCENT PostSourceServerBodyPlatform
	VMWARE  PostSourceServerBodyPlatform
	HYPERV  PostSourceServerBodyPlatform
	OTHER   PostSourceServerBodyPlatform
	DEFAULT PostSourceServerBodyPlatform
}

func GetPostSourceServerBodyPlatformEnum() PostSourceServerBodyPlatformEnum {
	return PostSourceServerBodyPlatformEnum{
		HW: PostSourceServerBodyPlatform{
			value: "hw",
		},
		ALI: PostSourceServerBodyPlatform{
			value: "ali",
		},
		AWS: PostSourceServerBodyPlatform{
			value: "aws",
		},
		AZURE: PostSourceServerBodyPlatform{
			value: "azure",
		},
		GCP: PostSourceServerBodyPlatform{
			value: "gcp",
		},
		TENCENT: PostSourceServerBodyPlatform{
			value: "tencent",
		},
		VMWARE: PostSourceServerBodyPlatform{
			value: "vmware",
		},
		HYPERV: PostSourceServerBodyPlatform{
			value: "hyperv",
		},
		OTHER: PostSourceServerBodyPlatform{
			value: "other",
		},
		DEFAULT: PostSourceServerBodyPlatform{
			value: "default",
		},
	}
}

func (c PostSourceServerBodyPlatform) Value() string {
	return c.value
}

func (c PostSourceServerBodyPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyPlatform) UnmarshalJSON(b []byte) error {
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
