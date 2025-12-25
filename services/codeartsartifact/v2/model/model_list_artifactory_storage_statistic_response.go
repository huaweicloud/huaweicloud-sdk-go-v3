package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListArtifactoryStorageStatisticResponse Response Object
type ListArtifactoryStorageStatisticResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListArtifactoryStorageStatisticResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// 参数解释: 存储容量数据。 取值范围: 不涉及。
	Result         *[]StorageStatisticDo `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListArtifactoryStorageStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArtifactoryStorageStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListArtifactoryStorageStatisticResponse", string(data)}, " ")
}

type ListArtifactoryStorageStatisticResponseStatus struct {
	value string
}

type ListArtifactoryStorageStatisticResponseStatusEnum struct {
	SUCCESS ListArtifactoryStorageStatisticResponseStatus
	ERROR   ListArtifactoryStorageStatisticResponseStatus
}

func GetListArtifactoryStorageStatisticResponseStatusEnum() ListArtifactoryStorageStatisticResponseStatusEnum {
	return ListArtifactoryStorageStatisticResponseStatusEnum{
		SUCCESS: ListArtifactoryStorageStatisticResponseStatus{
			value: "success",
		},
		ERROR: ListArtifactoryStorageStatisticResponseStatus{
			value: "error",
		},
	}
}

func (c ListArtifactoryStorageStatisticResponseStatus) Value() string {
	return c.value
}

func (c ListArtifactoryStorageStatisticResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListArtifactoryStorageStatisticResponseStatus) UnmarshalJSON(b []byte) error {
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
