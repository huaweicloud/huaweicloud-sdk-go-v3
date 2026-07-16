package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DetachDynamicStorageResponse Response Object
type DetachDynamicStorageResponse struct {

	// **参数解释**：存储类型。可选值为OBS/OBSFS/EFS。 **取值范围**：不涉及。
	Category *DetachDynamicStorageResponseCategory `json:"category,omitempty"`

	// **参数解释**：动态挂载实例ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：在Notebook实例中挂载的路径。 **取值范围**：不涉及。
	MountPath *string `json:"mount_path,omitempty"`

	// **参数解释**：动态挂载状态。 **取值范围**：枚举类型，取值如下： - MOUNTING：挂载中 - MOUNT_FAILED：挂载失败 - MOUNTED：已挂载 - UNMOUNTING：卸载中 - UNMOUNT_FAILED：卸载失败 - UNMOUNTED：卸载完成
	Status *string `json:"status,omitempty"`

	// **参数解释**：存储路径。 **取值范围**：不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**：挂载失败原因，动态挂载状态为MOUNT_FAILED时返回。 **取值范围**：不涉及。
	FailureReason *string `json:"failure_reason,omitempty"`

	// **参数解释**：EFS存储实例ID。 **取值范围**：不涉及。
	EfsId *string `json:"efs_id,omitempty"`

	// **参数解释**：存储挂载类型。 **取值范围**：枚举类型，取值如下：  - STATIC:不支持在实例运行期间挂载以及卸载的存储 - DYNAMIC:支持在实例运行期间挂载以及卸载的存储
	MountType      *string `json:"mount_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachDynamicStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachDynamicStorageResponse struct{}"
	}

	return strings.Join([]string{"DetachDynamicStorageResponse", string(data)}, " ")
}

type DetachDynamicStorageResponseCategory struct {
	value string
}

type DetachDynamicStorageResponseCategoryEnum struct {
	OBS   DetachDynamicStorageResponseCategory
	OBSFS DetachDynamicStorageResponseCategory
	EFS   DetachDynamicStorageResponseCategory
}

func GetDetachDynamicStorageResponseCategoryEnum() DetachDynamicStorageResponseCategoryEnum {
	return DetachDynamicStorageResponseCategoryEnum{
		OBS: DetachDynamicStorageResponseCategory{
			value: "OBS",
		},
		OBSFS: DetachDynamicStorageResponseCategory{
			value: "OBSFS",
		},
		EFS: DetachDynamicStorageResponseCategory{
			value: "EFS",
		},
	}
}

func (c DetachDynamicStorageResponseCategory) Value() string {
	return c.value
}

func (c DetachDynamicStorageResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachDynamicStorageResponseCategory) UnmarshalJSON(b []byte) error {
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
