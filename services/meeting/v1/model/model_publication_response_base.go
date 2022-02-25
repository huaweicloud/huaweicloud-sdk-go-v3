package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 发布响应信息
type PublicationResponseBase struct {
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

	PublishStatus *PublicationResponseBasePublishStatus `json:"publishStatus,omitempty"`
}

func (o PublicationResponseBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationResponseBase struct{}"
	}

	return strings.Join([]string{"PublicationResponseBase", string(data)}, " ")
}

type PublicationResponseBasePublishStatus struct {
	value string
}

type PublicationResponseBasePublishStatusEnum struct {
	NOT_ONLINE      PublicationResponseBasePublishStatus
	PUBLISHING      PublicationResponseBasePublishStatus
	ALREADY_OFFLINE PublicationResponseBasePublishStatus
}

func GetPublicationResponseBasePublishStatusEnum() PublicationResponseBasePublishStatusEnum {
	return PublicationResponseBasePublishStatusEnum{
		NOT_ONLINE: PublicationResponseBasePublishStatus{
			value: "NOT_ONLINE",
		},
		PUBLISHING: PublicationResponseBasePublishStatus{
			value: "PUBLISHING",
		},
		ALREADY_OFFLINE: PublicationResponseBasePublishStatus{
			value: "ALREADY_OFFLINE",
		},
	}
}

func (c PublicationResponseBasePublishStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicationResponseBasePublishStatus) UnmarshalJSON(b []byte) error {
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
