package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeMountRequest **参数解释**：实例的扩展存储配置 **约束限制**：最大数量为30。
type VolumeMountRequest struct {

	// **参数解释**：notebook支持的扩展存储类型，详见[[开发环境中如何选择存储](https://support.huaweicloud.com/usermanual-standard-modelarts/devtool-modelarts_0004.html#section7)](tag:hc)[[开发环境中如何选择存储](https://support.huaweicloud.com/intl/zh-cn/usermanual-standard-modelarts/devtool-modelarts_0004.html#section6)](tag:hk)[《用户指南》的“开发环境中如何选择存储”章节](tag:fcs,fcs-super) **约束限制**：不涉及 **默认取值**：不涉及。 **取值范围**：枚举类型，取值如下： - EVS：云硬盘 - OBS：对象存储服务 - OBSFS：并行文件系统（PFS） - EFS：弹性文件服务（SFS Turbo）
	Category VolumeMountRequestCategory `json:"category"`

	// **参数解释**：资源所属 **参数约束**：不涉及。 **取值范围**：枚举类型，取值如下： - MANAGED：托管，即资源在服务上。 - DEDICATED：非托管，即资源在用户账号上，只有在category为EFS时支持。 **默认取值**：不涉及。
	Ownership VolumeMountRequestOwnership `json:"ownership"`

	// **参数解释**：EFS专属存储盘uri或OBS并行文件系统路径 - EFS：登录弹性文件服务控制台，在文件系统列表中，单击文件系统名称进入详情页。其中，“共享路径”即为此参数的参数值。 - OBS：并行文件系统命名格式为：obs://<桶名>/<目录路径>/。登录对象存储服务控制台，在并行文件系统列表中，文件系统名称为桶名。单击文件系统名称进入详情页，在文件栏选择特定目录后，单击右侧“更多/复制路径”，该路径即为目录路径。 **参数约束**：只有当category为EFS或OBS或OBSFS，同时ownership为DEDICATED时必填，最大长度1024字符
	Uri *string `json:"uri,omitempty"`

	// **参数解释**：EFS专属存储盘ID，参数值获取方式如下：登录弹性文件服务控制台，在文件系统列表中，单击文件系统名称进入详情页。其中，“ID”即为此参数的参数值。 **参数约束**：只有当category为EFS，同时ownership为DEDICATED时必填。必须符合 UUID 格式（如 280a8bd5-03e2-4a5c-bea1-83d81e75bc53）。
	Id *string `json:"id,omitempty"`

	// **参数解释**：在Notebook实例中挂载的路径 **参数约束**：最大长度 256 字符
	MountPath *string `json:"mount_path,omitempty"`

	// **参数解释**：扩展存储挂载目录是否只读。默认值为false，可读写 **参数约束**：不涉及
	ReadOnly *bool `json:"read_only,omitempty"`

	// **参数解释**：DEW存储的用户AKSK凭据名称 **参数约束**：当category为OBS时必填，仅支持大小写字母、数字、中划线、下划线，长度 1-64 字符
	DewSecretName *string `json:"dew_secret_name,omitempty"`

	// **参数解释**：EVS云硬盘存储容量，单位GB。 **约束限制**：category为EVS时有效。 **取值范围**：不涉及。 **默认取值**：5。
	Capacity *int32 `json:"capacity,omitempty"`
}

func (o VolumeMountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeMountRequest struct{}"
	}

	return strings.Join([]string{"VolumeMountRequest", string(data)}, " ")
}

type VolumeMountRequestCategory struct {
	value string
}

type VolumeMountRequestCategoryEnum struct {
	OBS   VolumeMountRequestCategory
	OBSFS VolumeMountRequestCategory
	EFS   VolumeMountRequestCategory
}

func GetVolumeMountRequestCategoryEnum() VolumeMountRequestCategoryEnum {
	return VolumeMountRequestCategoryEnum{
		OBS: VolumeMountRequestCategory{
			value: "OBS",
		},
		OBSFS: VolumeMountRequestCategory{
			value: "OBSFS",
		},
		EFS: VolumeMountRequestCategory{
			value: "EFS",
		},
	}
}

func (c VolumeMountRequestCategory) Value() string {
	return c.value
}

func (c VolumeMountRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeMountRequestCategory) UnmarshalJSON(b []byte) error {
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

type VolumeMountRequestOwnership struct {
	value string
}

type VolumeMountRequestOwnershipEnum struct {
	MANAGED   VolumeMountRequestOwnership
	DEDICATED VolumeMountRequestOwnership
}

func GetVolumeMountRequestOwnershipEnum() VolumeMountRequestOwnershipEnum {
	return VolumeMountRequestOwnershipEnum{
		MANAGED: VolumeMountRequestOwnership{
			value: "MANAGED",
		},
		DEDICATED: VolumeMountRequestOwnership{
			value: "DEDICATED",
		},
	}
}

func (c VolumeMountRequestOwnership) Value() string {
	return c.value
}

func (c VolumeMountRequestOwnership) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeMountRequestOwnership) UnmarshalJSON(b []byte) error {
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
