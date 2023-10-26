package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteBackendTargetResponse Response Object
type DeleteBackendTargetResponse struct {

	// 后端存储库 id
	TargetId *string `json:"target_id,omitempty"`

	// 删除后端存储库时是否同时删除文件系统对应路径的数据
	DeleteDataInFileSystem *bool `json:"delete_data_in_file_system,omitempty"`

	// 后端存储库生命周期描述信息
	Lifecycle *DeleteBackendTargetResponseLifecycle `json:"lifecycle,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBackendTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendTargetResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackendTargetResponse", string(data)}, " ")
}

type DeleteBackendTargetResponseLifecycle struct {
	value string
}

type DeleteBackendTargetResponseLifecycleEnum struct {
	DELETING      DeleteBackendTargetResponseLifecycle
	AVAILABLE     DeleteBackendTargetResponseLifecycle
	MISCONFIGURED DeleteBackendTargetResponseLifecycle
	CREATING      DeleteBackendTargetResponseLifecycle
}

func GetDeleteBackendTargetResponseLifecycleEnum() DeleteBackendTargetResponseLifecycleEnum {
	return DeleteBackendTargetResponseLifecycleEnum{
		DELETING: DeleteBackendTargetResponseLifecycle{
			value: "DELETING",
		},
		AVAILABLE: DeleteBackendTargetResponseLifecycle{
			value: "AVAILABLE",
		},
		MISCONFIGURED: DeleteBackendTargetResponseLifecycle{
			value: "MISCONFIGURED",
		},
		CREATING: DeleteBackendTargetResponseLifecycle{
			value: "CREATING",
		},
	}
}

func (c DeleteBackendTargetResponseLifecycle) Value() string {
	return c.value
}

func (c DeleteBackendTargetResponseLifecycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteBackendTargetResponseLifecycle) UnmarshalJSON(b []byte) error {
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
