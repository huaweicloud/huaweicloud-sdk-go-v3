package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeRes **参数解释**：实例存储信息。
type VolumeRes struct {

	// **参数解释**：存储容量。 **取值范围**：EVS默认5G，EFS默认50G，最大限制4096G。
	Capacity *int32 `json:"capacity,omitempty"`

	// **参数解释**：支持的存储类型。不同存储类型的差异，详见[[开发环境中如何选择存储](https://support.huaweicloud.com/usermanual-standard-modelarts/devtool-modelarts_0004.html#section6)](tag:hc)[[开发环境中如何选择存储](https://support.huaweicloud.com/intl/zh-cn/usermanual-standard-modelarts/devtool-modelarts_0004.html#section5)](tag:hk)[《用户指南》的“开发环境中如何选择存储”章节](tag:fcs,fcs-super)。 **取值范围**：枚举类型，取值如下： - SFS：弹性文件服务 - EVS：云硬盘 - OBS：对象存储服务 - OBSFS：并行文件系统 - EFS：弹性文件服务（SFS Turbo）
	Category *string `json:"category,omitempty"`

	// **参数解释**：存储挂载至Notebook实例的目录，当前固定在/home/ma-user/work/下。 **取值范围**：不涉及。
	MountPath *string `json:"mount_path,omitempty"`

	// **参数解释**：资源所属。 **取值范围**：枚举类型，取值如下： - MANAGED：托管，即资源在服务上。 - DEDICATED：非托管，即资源在用户账号上，只有在category为EFS时支持。
	Ownership *string `json:"ownership,omitempty"`

	// **参数解释**：EVS扩容状态，扩容时的状态为RESIZING，此时实例可以正常使用。 **取值范围**：不涉及。
	Status *VolumeResStatus `json:"status,omitempty"`

	// **参数解释**：EFS专属存储盘ID或OBS存储ID，只有作为扩展存储时返回。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：扩展存储挂载目录是否只读。 **取值范围**：不涉及。
	ReadOnly *bool `json:"read_only,omitempty"`

	// **参数解释**：DEW存储的用户AKSK凭据名称。 **取值范围**：不涉及。
	DewSecretName *string `json:"dew_secret_name,omitempty"`

	// **参数解释**：规格包含的evs时，evs存储的sku编码。 **取值范围**：不涉及。
	EvsSkuCode *string `json:"evs_sku_code,omitempty"`

	// **参数解释**：只有当category为EFS或OBS或OBSFS时，挂载存储源路径。 **取值范围**：不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**：存储挂载类型。 **取值范围**：枚举类型，取值如下：  - STATIC:不支持在实例运行期间挂载以及卸载的存储 - DYNAMIC:支持在实例运行期间挂载以及卸载的存储
	MountType *string `json:"mount_type,omitempty"`
}

func (o VolumeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeRes struct{}"
	}

	return strings.Join([]string{"VolumeRes", string(data)}, " ")
}

type VolumeResStatus struct {
	value string
}

type VolumeResStatusEnum struct {
	DELETED        VolumeResStatus
	DELETE_FAILED  VolumeResStatus
	DELETING       VolumeResStatus
	IN_USE         VolumeResStatus
	MOUNTED        VolumeResStatus
	MOUNTING       VolumeResStatus
	MOUNT_FAILED   VolumeResStatus
	RESIZING       VolumeResStatus
	UNMOUNTED      VolumeResStatus
	UNMOUNTING     VolumeResStatus
	UNMOUNT_FAILED VolumeResStatus
}

func GetVolumeResStatusEnum() VolumeResStatusEnum {
	return VolumeResStatusEnum{
		DELETED: VolumeResStatus{
			value: "DELETED",
		},
		DELETE_FAILED: VolumeResStatus{
			value: "DELETE_FAILED",
		},
		DELETING: VolumeResStatus{
			value: "DELETING",
		},
		IN_USE: VolumeResStatus{
			value: "IN_USE",
		},
		MOUNTED: VolumeResStatus{
			value: "MOUNTED",
		},
		MOUNTING: VolumeResStatus{
			value: "MOUNTING",
		},
		MOUNT_FAILED: VolumeResStatus{
			value: "MOUNT_FAILED",
		},
		RESIZING: VolumeResStatus{
			value: "RESIZING",
		},
		UNMOUNTED: VolumeResStatus{
			value: "UNMOUNTED",
		},
		UNMOUNTING: VolumeResStatus{
			value: "UNMOUNTING",
		},
		UNMOUNT_FAILED: VolumeResStatus{
			value: "UNMOUNT_FAILED",
		},
	}
}

func (c VolumeResStatus) Value() string {
	return c.value
}

func (c VolumeResStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeResStatus) UnmarshalJSON(b []byte) error {
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
