package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackendTargetInfoResponse Response Object
type ShowBackendTargetInfoResponse struct {

	// 后端存储库 id
	TargetId *string `json:"target_id,omitempty"`

	// 后端存储库创建时间
	CreationTime *string `json:"creation_time,omitempty"`

	// 文件系统路径
	FileSystemPath *string `json:"file_system_path,omitempty"`

	FailureDetails *FailureDetailsMessage `json:"failure_details,omitempty"`

	// 后端存储库生命周期描述信息
	Lifecycle *ShowBackendTargetInfoResponseLifecycle `json:"lifecycle,omitempty"`

	Obs *ObsDataRepository `json:"obs,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackendTargetInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackendTargetInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBackendTargetInfoResponse", string(data)}, " ")
}

type ShowBackendTargetInfoResponseLifecycle struct {
	value string
}

type ShowBackendTargetInfoResponseLifecycleEnum struct {
	CREATING      ShowBackendTargetInfoResponseLifecycle
	AVAILABLE     ShowBackendTargetInfoResponseLifecycle
	MISCONFIGURED ShowBackendTargetInfoResponseLifecycle
	DELETING      ShowBackendTargetInfoResponseLifecycle
	FAILED        ShowBackendTargetInfoResponseLifecycle
}

func GetShowBackendTargetInfoResponseLifecycleEnum() ShowBackendTargetInfoResponseLifecycleEnum {
	return ShowBackendTargetInfoResponseLifecycleEnum{
		CREATING: ShowBackendTargetInfoResponseLifecycle{
			value: "CREATING",
		},
		AVAILABLE: ShowBackendTargetInfoResponseLifecycle{
			value: "AVAILABLE",
		},
		MISCONFIGURED: ShowBackendTargetInfoResponseLifecycle{
			value: "MISCONFIGURED",
		},
		DELETING: ShowBackendTargetInfoResponseLifecycle{
			value: "DELETING",
		},
		FAILED: ShowBackendTargetInfoResponseLifecycle{
			value: "FAILED",
		},
	}
}

func (c ShowBackendTargetInfoResponseLifecycle) Value() string {
	return c.value
}

func (c ShowBackendTargetInfoResponseLifecycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackendTargetInfoResponseLifecycle) UnmarshalJSON(b []byte) error {
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
