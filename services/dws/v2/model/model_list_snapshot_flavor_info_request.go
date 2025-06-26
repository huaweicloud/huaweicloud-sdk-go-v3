package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSnapshotFlavorInfoRequest Request Object
type ListSnapshotFlavorInfoRequest struct {

	// **参数解释**： 快照ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SnapshotId string `json:"snapshot_id"`

	// **参数解释**： 过滤快照规格类型，用于区分仅查询快照原始规格，还是包含可恢复的规格。 **约束限制**： 不涉及。 **取值范围**： snapshot：仅查询快照的规格信息 restore：同时查询恢复快照可用的规格信息 **默认取值**： snapshot
	Type *ListSnapshotFlavorInfoRequestType `json:"type,omitempty"`

	// **参数解释**： 恢复时的目标可用区，用以过滤目标可用区下可恢复的规格。 恢复3az集群时需传递3个可用区编码，英文逗号分割（无空格）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 快照原始集群所在可用区。
	AzCode *string `json:"az_code,omitempty"`

	// **参数解释**： 是否是细粒度备份恢复，用以过滤恢复时的可用规格。 **约束限制**： 不涉及。 **取值范围**： true|false **默认取值**： false
	FineGrainedRestore *bool `json:"fine_grained_restore,omitempty"`
}

func (o ListSnapshotFlavorInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotFlavorInfoRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotFlavorInfoRequest", string(data)}, " ")
}

type ListSnapshotFlavorInfoRequestType struct {
	value string
}

type ListSnapshotFlavorInfoRequestTypeEnum struct {
	SNAPSHOT ListSnapshotFlavorInfoRequestType
	RESTORE  ListSnapshotFlavorInfoRequestType
}

func GetListSnapshotFlavorInfoRequestTypeEnum() ListSnapshotFlavorInfoRequestTypeEnum {
	return ListSnapshotFlavorInfoRequestTypeEnum{
		SNAPSHOT: ListSnapshotFlavorInfoRequestType{
			value: "snapshot",
		},
		RESTORE: ListSnapshotFlavorInfoRequestType{
			value: "restore",
		},
	}
}

func (c ListSnapshotFlavorInfoRequestType) Value() string {
	return c.value
}

func (c ListSnapshotFlavorInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSnapshotFlavorInfoRequestType) UnmarshalJSON(b []byte) error {
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
