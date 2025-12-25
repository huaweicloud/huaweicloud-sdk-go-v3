package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLatestVersionFilesResponse Response Object
type ListLatestVersionFilesResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListLatestVersionFilesResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 最新版本列表。 **取值范围**: 不涉及。
	Result         *[]VersionViewVoV5 `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListLatestVersionFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatestVersionFilesResponse struct{}"
	}

	return strings.Join([]string{"ListLatestVersionFilesResponse", string(data)}, " ")
}

type ListLatestVersionFilesResponseStatus struct {
	value string
}

type ListLatestVersionFilesResponseStatusEnum struct {
	SUCCESS ListLatestVersionFilesResponseStatus
	ERROR   ListLatestVersionFilesResponseStatus
}

func GetListLatestVersionFilesResponseStatusEnum() ListLatestVersionFilesResponseStatusEnum {
	return ListLatestVersionFilesResponseStatusEnum{
		SUCCESS: ListLatestVersionFilesResponseStatus{
			value: "success",
		},
		ERROR: ListLatestVersionFilesResponseStatus{
			value: "error",
		},
	}
}

func (c ListLatestVersionFilesResponseStatus) Value() string {
	return c.value
}

func (c ListLatestVersionFilesResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLatestVersionFilesResponseStatus) UnmarshalJSON(b []byte) error {
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
