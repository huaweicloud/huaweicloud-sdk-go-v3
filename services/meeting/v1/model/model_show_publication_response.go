package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowPublicationResponse struct {
	// 发布ID

	Id *string `json:"id,omitempty"`
	// 更新者

	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// 创建时间

	CreateTime *int64 `json:"createTime,omitempty"`
	// 更新时间

	UpdateTime *int64 `json:"updateTime,omitempty"`
	// 发布名称

	PublishName *string `json:"publishName,omitempty"`
	// 发布范围

	PublishScope *string `json:"publishScope,omitempty"`
	// 开始时间

	StartTime *int64 `json:"startTime,omitempty"`
	// 结束时间

	EndTime *int64 `json:"endTime,omitempty"`
	// 根据当前时间确定发布状态 - NOT_ONLINE-未上线 - PUBLISHING-发布中 - ALREADY_OFFLINE-已下线

	PublishStatus *ShowPublicationResponsePublishStatus `json:"publishStatus,omitempty"`
	// 发布节目ID列表

	ProgramList *[]ProgramResponseBase `json:"programList,omitempty"`
	// 发布部门列表

	DeptList *[]PublishDeptResponseDto `json:"deptList,omitempty"`
	// 发布设备列表

	DeviceList     *[]PublishDeviceResponseDto `json:"deviceList,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowPublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicationResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicationResponse", string(data)}, " ")
}

type ShowPublicationResponsePublishStatus struct {
	value string
}

type ShowPublicationResponsePublishStatusEnum struct {
	NOT_ONLINE      ShowPublicationResponsePublishStatus
	PUBLISHING      ShowPublicationResponsePublishStatus
	ALREADY_OFFLINE ShowPublicationResponsePublishStatus
}

func GetShowPublicationResponsePublishStatusEnum() ShowPublicationResponsePublishStatusEnum {
	return ShowPublicationResponsePublishStatusEnum{
		NOT_ONLINE: ShowPublicationResponsePublishStatus{
			value: "NOT_ONLINE",
		},
		PUBLISHING: ShowPublicationResponsePublishStatus{
			value: "PUBLISHING",
		},
		ALREADY_OFFLINE: ShowPublicationResponsePublishStatus{
			value: "ALREADY_OFFLINE",
		},
	}
}

func (c ShowPublicationResponsePublishStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPublicationResponsePublishStatus) UnmarshalJSON(b []byte) error {
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
