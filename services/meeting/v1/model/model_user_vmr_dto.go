package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserVmrDto struct {

	// 云会议室的ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口中的vmrID。
	Id *string `json:"id,omitempty"`

	// 云会议室的固定会议ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口返回数据的vmrConferenceID。
	VmrId *string `json:"vmrId,omitempty"`

	// 云会议室名称。
	VmrName *string `json:"vmrName,omitempty"`

	// VMR模式。 - 0：个人会议ID - 1: 云会议室 - 2: 网络研讨会
	VmrMode *int32 `json:"vmrMode,omitempty"`

	// 云会议室套餐包的id，仅云会议室返回。
	VmrPkgId *string `json:"vmrPkgId,omitempty"`

	// 云会议室套餐包的名称，仅云会议室返回。
	VmrPkgName *string `json:"vmrPkgName,omitempty"`

	// 云会议室套餐包的会议并发方数，仅云会议室返回。
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty"`

	// 云会议室套餐包的与会时间，若为0则代表无限时长，仅云会议室返回。
	VmrPkgLength *int32 `json:"vmrPkgLength,omitempty"`

	// 云会议室状态。 * 0：正常 * 1：停用 * 2：未分配
	Status *UserVmrDtoStatus `json:"status,omitempty"`
}

func (o UserVmrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserVmrDto struct{}"
	}

	return strings.Join([]string{"UserVmrDto", string(data)}, " ")
}

type UserVmrDtoStatus struct {
	value int32
}

type UserVmrDtoStatusEnum struct {
	E_0 UserVmrDtoStatus
	E_1 UserVmrDtoStatus
	E_2 UserVmrDtoStatus
}

func GetUserVmrDtoStatusEnum() UserVmrDtoStatusEnum {
	return UserVmrDtoStatusEnum{
		E_0: UserVmrDtoStatus{
			value: 0,
		}, E_1: UserVmrDtoStatus{
			value: 1,
		}, E_2: UserVmrDtoStatus{
			value: 2,
		},
	}
}

func (c UserVmrDtoStatus) Value() int32 {
	return c.value
}

func (c UserVmrDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserVmrDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
