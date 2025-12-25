package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDomainReleaseRepoStorageResponse Response Object
type ShowDomainReleaseRepoStorageResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowDomainReleaseRepoStorageResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *ReleaseStorageVov5 `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowDomainReleaseRepoStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainReleaseRepoStorageResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainReleaseRepoStorageResponse", string(data)}, " ")
}

type ShowDomainReleaseRepoStorageResponseStatus struct {
	value string
}

type ShowDomainReleaseRepoStorageResponseStatusEnum struct {
	SUCCESS ShowDomainReleaseRepoStorageResponseStatus
	ERROR   ShowDomainReleaseRepoStorageResponseStatus
}

func GetShowDomainReleaseRepoStorageResponseStatusEnum() ShowDomainReleaseRepoStorageResponseStatusEnum {
	return ShowDomainReleaseRepoStorageResponseStatusEnum{
		SUCCESS: ShowDomainReleaseRepoStorageResponseStatus{
			value: "success",
		},
		ERROR: ShowDomainReleaseRepoStorageResponseStatus{
			value: "error",
		},
	}
}

func (c ShowDomainReleaseRepoStorageResponseStatus) Value() string {
	return c.value
}

func (c ShowDomainReleaseRepoStorageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainReleaseRepoStorageResponseStatus) UnmarshalJSON(b []byte) error {
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
