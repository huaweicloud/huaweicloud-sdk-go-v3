package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackendTargetInfoResponseBody 查询文件系统后端存储详情返回体
type ShowBackendTargetInfoResponseBody struct {

	// 绑定关系id
	TargetId *string `json:"target_id,omitempty"`

	// 绑定关系创建时间
	CreationTime *string `json:"creation_time,omitempty"`

	// 联动目录名称
	FileSystemPath *string `json:"file_system_path,omitempty"`

	FailureDetails *FailureDetailsMessage `json:"failure_details,omitempty"`

	// 绑定状态
	Lifecycle *ShowBackendTargetInfoResponseBodyLifecycle `json:"lifecycle,omitempty"`

	Obs *ObsDataRepository `json:"obs,omitempty"`
}

func (o ShowBackendTargetInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackendTargetInfoResponseBody struct{}"
	}

	return strings.Join([]string{"ShowBackendTargetInfoResponseBody", string(data)}, " ")
}

type ShowBackendTargetInfoResponseBodyLifecycle struct {
	value string
}

type ShowBackendTargetInfoResponseBodyLifecycleEnum struct {
	CREATING      ShowBackendTargetInfoResponseBodyLifecycle
	AVAILABLE     ShowBackendTargetInfoResponseBodyLifecycle
	MISCONFIGURED ShowBackendTargetInfoResponseBodyLifecycle
	DELETING      ShowBackendTargetInfoResponseBodyLifecycle
	FAILED        ShowBackendTargetInfoResponseBodyLifecycle
}

func GetShowBackendTargetInfoResponseBodyLifecycleEnum() ShowBackendTargetInfoResponseBodyLifecycleEnum {
	return ShowBackendTargetInfoResponseBodyLifecycleEnum{
		CREATING: ShowBackendTargetInfoResponseBodyLifecycle{
			value: "CREATING",
		},
		AVAILABLE: ShowBackendTargetInfoResponseBodyLifecycle{
			value: "AVAILABLE",
		},
		MISCONFIGURED: ShowBackendTargetInfoResponseBodyLifecycle{
			value: "MISCONFIGURED",
		},
		DELETING: ShowBackendTargetInfoResponseBodyLifecycle{
			value: "DELETING",
		},
		FAILED: ShowBackendTargetInfoResponseBodyLifecycle{
			value: "FAILED",
		},
	}
}

func (c ShowBackendTargetInfoResponseBodyLifecycle) Value() string {
	return c.value
}

func (c ShowBackendTargetInfoResponseBodyLifecycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackendTargetInfoResponseBodyLifecycle) UnmarshalJSON(b []byte) error {
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
