package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListElasticResourcePoolScaleRecordsRequest Request Object
type ListElasticResourcePoolScaleRecordsRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// start_time用于查询扩缩容历史的开始时间，该时间点需大于当前时间点减30天，必须小于end_time 。时间格式为unix时间戳，单位：毫秒。 ①若start_time为空，则查询end_time前七天到end_time的数据（end_time最大不能大于当前时间30天）。 ②查询当前时间点前15天到当前时间点的数据（start_time和end_time同时为空）。
	StartTime *int64 `json:"start_time,omitempty"`

	// end_time用于查询扩缩容历史的结束时间，该时间点不能小于开始时间，不能大于当前时间。时间格式为unix时间戳，单位：毫秒。 ①若end_time为空，则查询start_time到当前时间点的数据。 ②查询当前时间点前15天到当前时间的数据（start_time和end_time同时为空）。
	EndTime *int64 `json:"end_time,omitempty"`

	// 弹性资源池扩缩容的状态
	Status *ListElasticResourcePoolScaleRecordsRequestStatus `json:"status,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 当前分页条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListElasticResourcePoolScaleRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolScaleRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolScaleRecordsRequest", string(data)}, " ")
}

type ListElasticResourcePoolScaleRecordsRequestStatus struct {
	value string
}

type ListElasticResourcePoolScaleRecordsRequestStatusEnum struct {
	SUCCESS ListElasticResourcePoolScaleRecordsRequestStatus
	FAIL    ListElasticResourcePoolScaleRecordsRequestStatus
}

func GetListElasticResourcePoolScaleRecordsRequestStatusEnum() ListElasticResourcePoolScaleRecordsRequestStatusEnum {
	return ListElasticResourcePoolScaleRecordsRequestStatusEnum{
		SUCCESS: ListElasticResourcePoolScaleRecordsRequestStatus{
			value: "success",
		},
		FAIL: ListElasticResourcePoolScaleRecordsRequestStatus{
			value: "fail",
		},
	}
}

func (c ListElasticResourcePoolScaleRecordsRequestStatus) Value() string {
	return c.value
}

func (c ListElasticResourcePoolScaleRecordsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListElasticResourcePoolScaleRecordsRequestStatus) UnmarshalJSON(b []byte) error {
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
