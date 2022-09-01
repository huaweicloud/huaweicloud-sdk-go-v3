package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserVmrDto struct {

	// 云会议室的唯一标识
	Id *string `json:"id,omitempty" xml:"id"`

	// 云会议室ID即预约会议的id，分为两种，一种是创建用户时生成的，为用户个人的，另一种是企业管理员分配的专用云会议室
	VmrId *string `json:"vmrId,omitempty" xml:"vmrId"`

	// 云会议室名称
	VmrName *string `json:"vmrName,omitempty" xml:"vmrName"`

	// VMR模式。 - 0：个人会议ID - 1: 云会议室 - 2: 网络研讨会
	VmrMode *int32 `json:"vmrMode,omitempty" xml:"vmrMode"`

	// 云会议室套餐包的id，仅专用云会议室返回
	VmrPkgId *string `json:"vmrPkgId,omitempty" xml:"vmrPkgId"`

	// 云会议室套餐包的名称，仅专用云会议室返回
	VmrPkgName *string `json:"vmrPkgName,omitempty" xml:"vmrPkgName"`

	// 云会议室套餐包的会议并发方数，仅专用云会议室返回
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty" xml:"vmrPkgParties"`

	// 云会议室套餐包的与会时间，若为0则代表无限时长，仅专用云会议室返回
	VmrPkgLength *int32 `json:"vmrPkgLength,omitempty" xml:"vmrPkgLength"`

	// 云会议室状态。 * 0.正常 * 1.停用 * 2.未分配
	Status *UserVmrDtoStatus `json:"status,omitempty" xml:"status"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
