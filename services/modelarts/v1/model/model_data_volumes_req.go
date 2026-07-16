package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataVolumesReq 动态挂载存储请求体。
type DataVolumesReq struct {

	// **参数解释**：动态挂载存储类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - OBS：对象存储服务 - OBSFS：并行文件系统PFS - EFS：高性能弹性文件服务SFS Turbo  **默认取值**：无。
	Category DataVolumesReqCategory `json:"category"`

	// **参数解释**：在Notebook实例中挂载的路径。 **约束限制**：必须是/data目录的子目录。 **取值范围**：限制长度为256个字符，必须在Notebook的/data/的子目录下。 **默认取值**：无。
	MountPath string `json:"mount_path"`

	// **参数解释**：存储路径，示例：obs://modelarts/notebook/ 或 da669f6e-5591-4c10-b2a7-18d053a75677.sfsturbo.internal:/notebook。 **约束限制**：并行文件系统PFS 或 高性能弹性文件服务SFS Turbo中合法的挂载路径。 **取值范围**：限制长度为256个字符。 **默认取值**：不涉及。
	Uri string `json:"uri"`

	// **参数解释**：高性能弹性文件服务SFS Turbo实例id。 **约束限制**：若category字段为EFS，则此字段必填。 **取值范围**：合法UUID类型。 **默认取值**：无
	EfsId *string `json:"efs_id,omitempty"`
}

func (o DataVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataVolumesReq struct{}"
	}

	return strings.Join([]string{"DataVolumesReq", string(data)}, " ")
}

type DataVolumesReqCategory struct {
	value string
}

type DataVolumesReqCategoryEnum struct {
	OBS   DataVolumesReqCategory
	OBSFS DataVolumesReqCategory
	EFS   DataVolumesReqCategory
}

func GetDataVolumesReqCategoryEnum() DataVolumesReqCategoryEnum {
	return DataVolumesReqCategoryEnum{
		OBS: DataVolumesReqCategory{
			value: "OBS",
		},
		OBSFS: DataVolumesReqCategory{
			value: "OBSFS",
		},
		EFS: DataVolumesReqCategory{
			value: "EFS",
		},
	}
}

func (c DataVolumesReqCategory) Value() string {
	return c.value
}

func (c DataVolumesReqCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataVolumesReqCategory) UnmarshalJSON(b []byte) error {
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
