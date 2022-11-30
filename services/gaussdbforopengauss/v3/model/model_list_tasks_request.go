package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTasksRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 任务状态。
	Status *ListTasksRequestStatus `json:"status,omitempty"`

	// 任务名称。
	Name *ListTasksRequestName `json:"name,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestStatus struct {
	value string
}

type ListTasksRequestStatusEnum struct {
	RUNNING   ListTasksRequestStatus
	COMPLETED ListTasksRequestStatus
	FAILED    ListTasksRequestStatus
}

func GetListTasksRequestStatusEnum() ListTasksRequestStatusEnum {
	return ListTasksRequestStatusEnum{
		RUNNING: ListTasksRequestStatus{
			value: "Running",
		},
		COMPLETED: ListTasksRequestStatus{
			value: "Completed",
		},
		FAILED: ListTasksRequestStatus{
			value: "Failed",
		},
	}
}

func (c ListTasksRequestStatus) Value() string {
	return c.value
}

func (c ListTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListTasksRequestName struct {
	value string
}

type ListTasksRequestNameEnum struct {
	CREATE_GAUSS_DBV5_INSTANCE                     ListTasksRequestName
	BACKUP_SNAPSHOT_GAUSS_DBV5_IN_INSTANCE         ListTasksRequestName
	CLONE_GAUSS_DBV5_NEW_INSTANCE                  ListTasksRequestName
	RESTORE_GAUSS_DBV5_IN_INSTANCE                 ListTasksRequestName
	RESTORE_GAUSS_DBV5_IN_INSTANCE_TO_EXISTED_INST ListTasksRequestName
	DELETE_GAUSS_DBV5_INSTANCE                     ListTasksRequestName
	ENLARGE_GAUSS_DBV5_VOLUME                      ListTasksRequestName
	RESIZE_GAUSS_DBV5_FLAVOR                       ListTasksRequestName
	GAUSS_DBV5_EXPAND_CLUSTER_CN                   ListTasksRequestName
	GAUSS_DBV5_EXPAND_CLUSTER_DN                   ListTasksRequestName
}

func GetListTasksRequestNameEnum() ListTasksRequestNameEnum {
	return ListTasksRequestNameEnum{
		CREATE_GAUSS_DBV5_INSTANCE: ListTasksRequestName{
			value: "CreateGaussDBV5Instance",
		},
		BACKUP_SNAPSHOT_GAUSS_DBV5_IN_INSTANCE: ListTasksRequestName{
			value: "BackupSnapshotGaussDBV5InInstance",
		},
		CLONE_GAUSS_DBV5_NEW_INSTANCE: ListTasksRequestName{
			value: "CloneGaussDBV5NewInstance",
		},
		RESTORE_GAUSS_DBV5_IN_INSTANCE: ListTasksRequestName{
			value: "RestoreGaussDBV5InInstance",
		},
		RESTORE_GAUSS_DBV5_IN_INSTANCE_TO_EXISTED_INST: ListTasksRequestName{
			value: "RestoreGaussDBV5InInstanceToExistedInst",
		},
		DELETE_GAUSS_DBV5_INSTANCE: ListTasksRequestName{
			value: "DeleteGaussDBV5Instance",
		},
		ENLARGE_GAUSS_DBV5_VOLUME: ListTasksRequestName{
			value: "EnlargeGaussDBV5Volume",
		},
		RESIZE_GAUSS_DBV5_FLAVOR: ListTasksRequestName{
			value: "ResizeGaussDBV5Flavor",
		},
		GAUSS_DBV5_EXPAND_CLUSTER_CN: ListTasksRequestName{
			value: "GaussDBV5ExpandClusterCN",
		},
		GAUSS_DBV5_EXPAND_CLUSTER_DN: ListTasksRequestName{
			value: "GaussDBV5ExpandClusterDN",
		},
	}
}

func (c ListTasksRequestName) Value() string {
	return c.value
}

func (c ListTasksRequestName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestName) UnmarshalJSON(b []byte) error {
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
