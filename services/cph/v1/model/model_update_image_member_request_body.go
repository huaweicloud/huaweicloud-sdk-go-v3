package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateImageMemberRequestBody 更新共享镜像成员请求体
type UpdateImageMemberRequestBody struct {

	// 接受状态。   - accepted: 接受 - rejected: 拒绝
	Status UpdateImageMemberRequestBodyStatus `json:"status"`
}

func (o UpdateImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateImageMemberRequestBody", string(data)}, " ")
}

type UpdateImageMemberRequestBodyStatus struct {
	value string
}

type UpdateImageMemberRequestBodyStatusEnum struct {
	ACCEPTED UpdateImageMemberRequestBodyStatus
	REJECTED UpdateImageMemberRequestBodyStatus
}

func GetUpdateImageMemberRequestBodyStatusEnum() UpdateImageMemberRequestBodyStatusEnum {
	return UpdateImageMemberRequestBodyStatusEnum{
		ACCEPTED: UpdateImageMemberRequestBodyStatus{
			value: "accepted",
		},
		REJECTED: UpdateImageMemberRequestBodyStatus{
			value: "rejected",
		},
	}
}

func (c UpdateImageMemberRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateImageMemberRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageMemberRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
