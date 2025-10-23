package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRepositoryResourcePermissionsResponse Response Object
type UpdateRepositoryResourcePermissionsResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *UpdateRepositoryResourcePermissionsResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                                `json:"-"`
}

func (o UpdateRepositoryResourcePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryResourcePermissionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryResourcePermissionsResponse", string(data)}, " ")
}

type UpdateRepositoryResourcePermissionsResponseStatus struct {
	value string
}

type UpdateRepositoryResourcePermissionsResponseStatusEnum struct {
	SUCCESS UpdateRepositoryResourcePermissionsResponseStatus
	FAIL    UpdateRepositoryResourcePermissionsResponseStatus
}

func GetUpdateRepositoryResourcePermissionsResponseStatusEnum() UpdateRepositoryResourcePermissionsResponseStatusEnum {
	return UpdateRepositoryResourcePermissionsResponseStatusEnum{
		SUCCESS: UpdateRepositoryResourcePermissionsResponseStatus{
			value: "success",
		},
		FAIL: UpdateRepositoryResourcePermissionsResponseStatus{
			value: "fail",
		},
	}
}

func (c UpdateRepositoryResourcePermissionsResponseStatus) Value() string {
	return c.value
}

func (c UpdateRepositoryResourcePermissionsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRepositoryResourcePermissionsResponseStatus) UnmarshalJSON(b []byte) error {
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
