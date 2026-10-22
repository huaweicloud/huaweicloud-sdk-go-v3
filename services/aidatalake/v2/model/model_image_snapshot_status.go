package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageSnapshotStatus **参数解释**：镜像快照状态。 **约束限制**：不涉及。 **取值范围**：   - Creating：创建中。   - Available：可用。   - Unavailable：不可用。 **默认取值**：不涉及。
type ImageSnapshotStatus struct {
	value string
}

type ImageSnapshotStatusEnum struct {
	CREATING    ImageSnapshotStatus
	AVAILABLE   ImageSnapshotStatus
	UNAVAILABLE ImageSnapshotStatus
}

func GetImageSnapshotStatusEnum() ImageSnapshotStatusEnum {
	return ImageSnapshotStatusEnum{
		CREATING: ImageSnapshotStatus{
			value: "Creating",
		},
		AVAILABLE: ImageSnapshotStatus{
			value: "Available",
		},
		UNAVAILABLE: ImageSnapshotStatus{
			value: "Unavailable",
		},
	}
}

func (c ImageSnapshotStatus) Value() string {
	return c.value
}

func (c ImageSnapshotStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSnapshotStatus) UnmarshalJSON(b []byte) error {
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
