package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListApiVersionsResponse struct {
	// API版本号。 取值：由高到低版本分别为v3，v2，v2.0。

	Id *string `json:"id,omitempty"`
	// API版本的状态。 取值： - CURRENT：当前版本。 - STABLE：稳定版本。 - DEPRECATED：废弃版本。

	Status         *ListApiVersionsResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}

type ListApiVersionsResponseStatus struct {
	value string
}

type ListApiVersionsResponseStatusEnum struct {
	CURRENT    ListApiVersionsResponseStatus
	STABLE     ListApiVersionsResponseStatus
	DEPRECATED ListApiVersionsResponseStatus
}

func GetListApiVersionsResponseStatusEnum() ListApiVersionsResponseStatusEnum {
	return ListApiVersionsResponseStatusEnum{
		CURRENT: ListApiVersionsResponseStatus{
			value: "CURRENT",
		},
		STABLE: ListApiVersionsResponseStatus{
			value: "STABLE",
		},
		DEPRECATED: ListApiVersionsResponseStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ListApiVersionsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiVersionsResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
