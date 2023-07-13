package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAcceleratorsRequest Request Object
type ListAcceleratorsRequest struct {

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 资源ID。
	Id *string `json:"id,omitempty"`

	// 资源名称，取值范围：0~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 配置状态，可选范围: - ACTIVE：运行中 - PENDING：待定 - ERROR：错误 - DELETING：正在删除
	Status *ListAcceleratorsRequestStatus `json:"status,omitempty"`

	// 租户的企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListAcceleratorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAcceleratorsRequest struct{}"
	}

	return strings.Join([]string{"ListAcceleratorsRequest", string(data)}, " ")
}

type ListAcceleratorsRequestStatus struct {
	value string
}

type ListAcceleratorsRequestStatusEnum struct {
	ACTIVE   ListAcceleratorsRequestStatus
	PENDING  ListAcceleratorsRequestStatus
	ERROR    ListAcceleratorsRequestStatus
	DELETING ListAcceleratorsRequestStatus
}

func GetListAcceleratorsRequestStatusEnum() ListAcceleratorsRequestStatusEnum {
	return ListAcceleratorsRequestStatusEnum{
		ACTIVE: ListAcceleratorsRequestStatus{
			value: "ACTIVE",
		},
		PENDING: ListAcceleratorsRequestStatus{
			value: "PENDING",
		},
		ERROR: ListAcceleratorsRequestStatus{
			value: "ERROR",
		},
		DELETING: ListAcceleratorsRequestStatus{
			value: "DELETING",
		},
	}
}

func (c ListAcceleratorsRequestStatus) Value() string {
	return c.value
}

func (c ListAcceleratorsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAcceleratorsRequestStatus) UnmarshalJSON(b []byte) error {
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
