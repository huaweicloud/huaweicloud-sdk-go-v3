package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDedicatedHostsRequest Request Object
type ListDedicatedHostsRequest struct {

	// 专属主机ID。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 专属主机名称。
	Name *string `json:"name,omitempty"`

	// 专属主机类型。
	HostType *string `json:"host_type,omitempty"`

	// 专属主机类型的名称。
	HostTypeName *string `json:"host_type_name,omitempty"`

	// 规格ID。
	Flavor *string `json:"flavor,omitempty"`

	// 专属主机状态。  取值范围：“available”、“fault”或“released”。
	State *ListDedicatedHostsRequestState `json:"state,omitempty"`

	// 取值范围：租户ID或“all”。  只有管理员可以指定该参数。
	Tenant *string `json:"tenant,omitempty"`

	// 专属主机所属AZ。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 每个页面上显示的条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 该值是上一页最后一条记录的ID。  如果“marker”取值无效，将会返回“400”错误码。
	Marker *string `json:"marker,omitempty"`

	// 专属主机标签。
	Tags *string `json:"tags,omitempty"`

	// 专属主机上的云服务器ID。
	InstanceUuid *string `json:"instance_uuid,omitempty"`

	// 专属主机的释放时间。
	ReleasedAt *string `json:"released_at,omitempty"`

	// 当专属主机更新了状态时，按日期和时间戳过滤响应。为了便于记录更改，还可能返回最近删除的专属主机。  日期和时间戳的格式为ISO 8601：CCYY-MM-DDThh:mm:ss±hh:mm  如果包含“hh:mm”值，则将时区作为UTC的偏移量返回。例如，“2015-08-27T09:49:58-05:00”。如果您省略时区，则假定为UTC时区。
	ChangesSince *string `json:"changes-since,omitempty"`
}

func (o ListDedicatedHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostsRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostsRequest", string(data)}, " ")
}

type ListDedicatedHostsRequestState struct {
	value string
}

type ListDedicatedHostsRequestStateEnum struct {
	AVAILABLE ListDedicatedHostsRequestState
	FAULT     ListDedicatedHostsRequestState
	RELEASED  ListDedicatedHostsRequestState
}

func GetListDedicatedHostsRequestStateEnum() ListDedicatedHostsRequestStateEnum {
	return ListDedicatedHostsRequestStateEnum{
		AVAILABLE: ListDedicatedHostsRequestState{
			value: "available",
		},
		FAULT: ListDedicatedHostsRequestState{
			value: "fault",
		},
		RELEASED: ListDedicatedHostsRequestState{
			value: "released",
		},
	}
}

func (c ListDedicatedHostsRequestState) Value() string {
	return c.value
}

func (c ListDedicatedHostsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDedicatedHostsRequestState) UnmarshalJSON(b []byte) error {
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
