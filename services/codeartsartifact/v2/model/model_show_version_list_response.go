package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowVersionListResponse Response Object
type ShowVersionListResponse struct {

	// **参数解释**： 查询发布库版本列表成功或失败的状态。 **取值范围**： - success：查询发布库版本列表成功。 - error：查询发布库版本列表失败。
	Status *ShowVersionListResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求的唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**： 版本信息列表。 **取值范围**： 不涉及。
	Result         *[]VersionListViewV5 `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowVersionListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionListResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionListResponse", string(data)}, " ")
}

type ShowVersionListResponseStatus struct {
	value string
}

type ShowVersionListResponseStatusEnum struct {
	SUCCESS ShowVersionListResponseStatus
	ERROR   ShowVersionListResponseStatus
}

func GetShowVersionListResponseStatusEnum() ShowVersionListResponseStatusEnum {
	return ShowVersionListResponseStatusEnum{
		SUCCESS: ShowVersionListResponseStatus{
			value: "success",
		},
		ERROR: ShowVersionListResponseStatus{
			value: "error",
		},
	}
}

func (c ShowVersionListResponseStatus) Value() string {
	return c.value
}

func (c ShowVersionListResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVersionListResponseStatus) UnmarshalJSON(b []byte) error {
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
