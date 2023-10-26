package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBackendTargetResponse Response Object
type CreateBackendTargetResponse struct {

	// 后端存储库 id据
	TargetId *string `json:"target_id,omitempty"`

	// 后端存储库创建时间
	CreationTime *string `json:"creation_time,omitempty"`

	FailureDetails *FailureDetailsMessage `json:"failure_details,omitempty"`

	// 文件系统路径
	FileSystemPath *string `json:"file_system_path,omitempty"`

	// 后端存储库声明周期描述信息
	Lifecycle *CreateBackendTargetResponseLifecycle `json:"lifecycle,omitempty"`

	Obs *ObsDataRepository `json:"obs,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackendTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackendTargetResponse struct{}"
	}

	return strings.Join([]string{"CreateBackendTargetResponse", string(data)}, " ")
}

type CreateBackendTargetResponseLifecycle struct {
	value string
}

type CreateBackendTargetResponseLifecycleEnum struct {
	CREATING      CreateBackendTargetResponseLifecycle
	AVAILABLE     CreateBackendTargetResponseLifecycle
	MISCONFIGURED CreateBackendTargetResponseLifecycle
	DELETING      CreateBackendTargetResponseLifecycle
}

func GetCreateBackendTargetResponseLifecycleEnum() CreateBackendTargetResponseLifecycleEnum {
	return CreateBackendTargetResponseLifecycleEnum{
		CREATING: CreateBackendTargetResponseLifecycle{
			value: "CREATING",
		},
		AVAILABLE: CreateBackendTargetResponseLifecycle{
			value: "AVAILABLE",
		},
		MISCONFIGURED: CreateBackendTargetResponseLifecycle{
			value: "MISCONFIGURED",
		},
		DELETING: CreateBackendTargetResponseLifecycle{
			value: "DELETING",
		},
	}
}

func (c CreateBackendTargetResponseLifecycle) Value() string {
	return c.value
}

func (c CreateBackendTargetResponseLifecycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBackendTargetResponseLifecycle) UnmarshalJSON(b []byte) error {
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
