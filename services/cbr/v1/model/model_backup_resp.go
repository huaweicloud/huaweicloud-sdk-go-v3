package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BackupResp struct {

	// 还原点ID
	CheckpointId string `json:"checkpoint_id"`

	// 创建时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 备份描述
	Description string `json:"description"`

	// 过期时间，例如:\"2020-02-05T10:38:34.209782\"
	ExpiredAt *sdktime.SdkTime `json:"expired_at"`

	ExtendInfo *BackupExtendInfo `json:"extend_info"`

	// 备份ID
	Id string `json:"id"`

	// [备份类型。取值为backup和replication。](tag:hws,hws_hk,ocb) [备份类型。取值为backup。](tag:g42,hk_g42,sbc,dt,fcs_vm,ctc,tm,tlf,cmcc,hcso_dt)
	ImageType string `json:"image_type"`

	// 备份名称
	Name string `json:"name"`

	// 父备份ID
	ParentId string `json:"parent_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 备份时间
	ProtectedAt string `json:"protected_at"`

	// 资源可用区
	ResourceAz string `json:"resource_az"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源大小，单位为GB
	ResourceSize int32 `json:"resource_size"`

	// [资源类型: OS::Nova::Server, OS::Cinder::Volume, OS::Ironic::BareMetalServer, OS::Native::Server, OS::Sfs::Turbo, OS::Workspace::DesktopV2](tag:hws,hws_hk) [资源类型: OS::Nova::Server, OS::Cinder::Volume, OS::Sfs::Turbo](tag:hk_g42,sbc,dt) [资源类型: OS::Nova::Server, OS::Cinder::Volume, OS::Ironic::BareMetalServer, OS::Sfs::Turbo](tag:fcs_vm,ctc,ocb,tm) [资源类型: OS::Nova::Server, OS::Cinder::Volume](tag:tlf,cmcc,hcso_dt) [资源类型: OS::Nova::Server, OS::Cinder::Volume, OS::Sfs::Turbo, OS::Workspace::DesktopV2](tag:g42)
	ResourceType string `json:"resource_type"`

	// 备份状态 - available: 可用 - protecting: 保护中 - deleting: 删除中 - restoring: 恢复中 - error: 异常 - waiting_protect: 等待保护 - waiting_delete: 等待删除 - waiting_restore: 等待恢复
	Status BackupRespStatus `json:"status"`

	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 存储库ID
	VaultId string `json:"vault_id"`

	// 复制记录
	ReplicationRecords *[]ReplicationRecordGet `json:"replication_records,omitempty"`

	// 企业项目id,默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 备份提供商ID，用于区分备份对象。当前取值包含： [0daac4c5-6707-4851-97ba-169e36266b66，该值代表备份对象为云服务器。d1603440-187d-4516-af25-121250c7cc97，该值代表备份对象为云硬盘。3f3c3220-245c-4805-b811-758870015881， 该值代表备份对象为SFS Turbo。a13639de-00be-4e94-af30-26912d75e4a2，该值代表备份对象为混合云VMware备份。](tag:hws,hws_hk) [0daac4c5-6707-4851-97ba-169e36266b66，该值代表备份对象为云服务器。d1603440-187d-4516-af25-121250c7cc97，该值代表备份对象为云硬盘。3f3c3220-245c-4805-b811-758870015881，该值代表备份对象为SFS Turbo。](tag:ocb,tlf,sbc,fcs_vm,g42,tm,dt,cmcc) [0daac4c5-6707-4851-97ba-169e36266b66，该值代表备份对象为云服务器。d1603440-187d-4516-af25-121250c7cc97，该值代表备份对象为云硬盘。3f3c3220-245c-4805-b811-758870015881，该值代表备份对象为SFS Turbo。86a80900-71bf-4961-956a-d52df944f84a，该值代表备份对象为Workspace。](tag:ctc) [0daac4c5-6707-4851-97ba-169e36266b66，该值代表备份对象为云服务器。d1603440-187d-4516-af25-121250c7cc97，该值代表备份对象为云硬盘。](tag:hcso_dt)
	ProviderId string `json:"provider_id"`

	// 子副本列表
	Children *[]BackupResp `json:"children,omitempty"`

	// 是否是增备
	Incremental *bool `json:"incremental,omitempty"`

	// 备份副本快照类型
	Version *int32 `json:"version,omitempty"`
}

func (o BackupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResp struct{}"
	}

	return strings.Join([]string{"BackupResp", string(data)}, " ")
}

type BackupRespStatus struct {
	value string
}

type BackupRespStatusEnum struct {
	AVAILABLE       BackupRespStatus
	PROTECTING      BackupRespStatus
	DELETING        BackupRespStatus
	RESTORING       BackupRespStatus
	ERROR           BackupRespStatus
	WAITING_PROTECT BackupRespStatus
	WAITING_DELETE  BackupRespStatus
	WAITING_RESTORE BackupRespStatus
}

func GetBackupRespStatusEnum() BackupRespStatusEnum {
	return BackupRespStatusEnum{
		AVAILABLE: BackupRespStatus{
			value: "available",
		},
		PROTECTING: BackupRespStatus{
			value: "protecting",
		},
		DELETING: BackupRespStatus{
			value: "deleting",
		},
		RESTORING: BackupRespStatus{
			value: "restoring",
		},
		ERROR: BackupRespStatus{
			value: "error",
		},
		WAITING_PROTECT: BackupRespStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: BackupRespStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: BackupRespStatus{
			value: "waiting_restore",
		},
	}
}

func (c BackupRespStatus) Value() string {
	return c.value
}

func (c BackupRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupRespStatus) UnmarshalJSON(b []byte) error {
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
