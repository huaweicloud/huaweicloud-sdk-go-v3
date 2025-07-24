package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteBackendTargetResponse Response Object
type DeleteBackendTargetResponse struct {

	// 绑定关系ID
	TargetId *string `json:"target_id,omitempty"`

	// 删除后端存储时是否同时删除文件系统内的联动目录及其数据文件
	DeleteDataInFileSystem *bool `json:"delete_data_in_file_system,omitempty"`

	// 绑定状态。只支持DELETING和FAILED
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
	DELETING DeleteBackendTargetResponseLifecycle
	FAILED   DeleteBackendTargetResponseLifecycle
}

func GetDeleteBackendTargetResponseLifecycleEnum() DeleteBackendTargetResponseLifecycleEnum {
	return DeleteBackendTargetResponseLifecycleEnum{
		DELETING: DeleteBackendTargetResponseLifecycle{
			value: "DELETING",
		},
		FAILED: DeleteBackendTargetResponseLifecycle{
			value: "FAILED",
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
