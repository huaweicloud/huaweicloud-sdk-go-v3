package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowAssetMetaRequest struct {
	// 媒资id。一次最多10个<br/>

	AssetId *[]string `json:"asset_id,omitempty"`
	// 媒资状态<br/>

	Status *[]ShowAssetMetaRequestStatus `json:"status,omitempty"`
	// 转码状态<br/>

	TranscodeStatus *[]ShowAssetMetaRequestTranscodeStatus `json:"transcodeStatus,omitempty"`
	// 媒资状态<br/>

	AssetStatus *[]ShowAssetMetaRequestAssetStatus `json:"assetStatus,omitempty"`
	// 起始时间.指定task_id时该参数无效<br/>

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间.指定task_id时该参数无效<br/>

	EndTime *string `json:"end_time,omitempty"`
	// 分类ID<br/>

	CategoryId *int32 `json:"category_id,omitempty"`
	// 视频标签。<br/>单个标签不超过16个字节，最多不超过16个标签。<br/> 多个用逗号分隔，UTF8编码。<br/>

	Tags *string `json:"tags,omitempty"`
	// 在媒资标题、描述中模糊查询的字符串。<br/>暂不支持模糊查询。<br/>

	QueryString *string `json:"query_string,omitempty"`
	// 分页编号,默认为0。<br/>

	Page *int32 `json:"page,omitempty"`
	// 每页记录数。默认10，范围[1,100]<br/> 指定task_id时该参数无效<br/>

	Size *int32 `json:"size,omitempty"`

	Order *ShowAssetMetaRequestOrder `json:"order,omitempty"`
}

func (o ShowAssetMetaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAssetMetaRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetMetaRequest", string(data)}, " ")
}

type ShowAssetMetaRequestStatus struct {
	value string
}

type ShowAssetMetaRequestStatusEnum struct {
	UNCREATED         ShowAssetMetaRequestStatus
	DELETED           ShowAssetMetaRequestStatus
	CANCELLED         ShowAssetMetaRequestStatus
	SERVER_ERROR      ShowAssetMetaRequestStatus
	UPLOAD_FAILED     ShowAssetMetaRequestStatus
	CREATING          ShowAssetMetaRequestStatus
	PUBLISHED         ShowAssetMetaRequestStatus
	WAITING_TRANSCODE ShowAssetMetaRequestStatus
	TRANSCODING       ShowAssetMetaRequestStatus
	TRANSCODE_FAILED  ShowAssetMetaRequestStatus
	TRANSCODE_SUCCEED ShowAssetMetaRequestStatus
	CREATED           ShowAssetMetaRequestStatus
}

func GetShowAssetMetaRequestStatusEnum() ShowAssetMetaRequestStatusEnum {
	return ShowAssetMetaRequestStatusEnum{
		UNCREATED: ShowAssetMetaRequestStatus{
			value: "UNCREATED",
		},
		DELETED: ShowAssetMetaRequestStatus{
			value: "DELETED",
		},
		CANCELLED: ShowAssetMetaRequestStatus{
			value: "CANCELLED",
		},
		SERVER_ERROR: ShowAssetMetaRequestStatus{
			value: "SERVER_ERROR",
		},
		UPLOAD_FAILED: ShowAssetMetaRequestStatus{
			value: "UPLOAD_FAILED",
		},
		CREATING: ShowAssetMetaRequestStatus{
			value: "CREATING",
		},
		PUBLISHED: ShowAssetMetaRequestStatus{
			value: "PUBLISHED",
		},
		WAITING_TRANSCODE: ShowAssetMetaRequestStatus{
			value: "WAITING_TRANSCODE",
		},
		TRANSCODING: ShowAssetMetaRequestStatus{
			value: "TRANSCODING",
		},
		TRANSCODE_FAILED: ShowAssetMetaRequestStatus{
			value: "TRANSCODE_FAILED",
		},
		TRANSCODE_SUCCEED: ShowAssetMetaRequestStatus{
			value: "TRANSCODE_SUCCEED",
		},
		CREATED: ShowAssetMetaRequestStatus{
			value: "CREATED",
		},
	}
}

func (c ShowAssetMetaRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowAssetMetaRequestStatus) UnmarshalJSON(b []byte) error {
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

type ShowAssetMetaRequestTranscodeStatus struct {
	value string
}

type ShowAssetMetaRequestTranscodeStatusEnum struct {
	TRANSCODING       ShowAssetMetaRequestTranscodeStatus
	TRANSCODE_FAILED  ShowAssetMetaRequestTranscodeStatus
	TRANSCODE_SUCCEED ShowAssetMetaRequestTranscodeStatus
	UN_TRANSCODE      ShowAssetMetaRequestTranscodeStatus
	WAITING_TRANSCODE ShowAssetMetaRequestTranscodeStatus
}

func GetShowAssetMetaRequestTranscodeStatusEnum() ShowAssetMetaRequestTranscodeStatusEnum {
	return ShowAssetMetaRequestTranscodeStatusEnum{
		TRANSCODING: ShowAssetMetaRequestTranscodeStatus{
			value: "TRANSCODING",
		},
		TRANSCODE_FAILED: ShowAssetMetaRequestTranscodeStatus{
			value: "TRANSCODE_FAILED",
		},
		TRANSCODE_SUCCEED: ShowAssetMetaRequestTranscodeStatus{
			value: "TRANSCODE_SUCCEED",
		},
		UN_TRANSCODE: ShowAssetMetaRequestTranscodeStatus{
			value: "UN_TRANSCODE",
		},
		WAITING_TRANSCODE: ShowAssetMetaRequestTranscodeStatus{
			value: "WAITING_TRANSCODE",
		},
	}
}

func (c ShowAssetMetaRequestTranscodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowAssetMetaRequestTranscodeStatus) UnmarshalJSON(b []byte) error {
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

type ShowAssetMetaRequestAssetStatus struct {
	value string
}

type ShowAssetMetaRequestAssetStatusEnum struct {
	PUBLISHED ShowAssetMetaRequestAssetStatus
	CREATED   ShowAssetMetaRequestAssetStatus
}

func GetShowAssetMetaRequestAssetStatusEnum() ShowAssetMetaRequestAssetStatusEnum {
	return ShowAssetMetaRequestAssetStatusEnum{
		PUBLISHED: ShowAssetMetaRequestAssetStatus{
			value: "PUBLISHED",
		},
		CREATED: ShowAssetMetaRequestAssetStatus{
			value: "CREATED",
		},
	}
}

func (c ShowAssetMetaRequestAssetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowAssetMetaRequestAssetStatus) UnmarshalJSON(b []byte) error {
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

type ShowAssetMetaRequestOrder struct {
	value string
}

type ShowAssetMetaRequestOrderEnum struct {
	ASC  ShowAssetMetaRequestOrder
	DESC ShowAssetMetaRequestOrder
}

func GetShowAssetMetaRequestOrderEnum() ShowAssetMetaRequestOrderEnum {
	return ShowAssetMetaRequestOrderEnum{
		ASC: ShowAssetMetaRequestOrder{
			value: "asc",
		},
		DESC: ShowAssetMetaRequestOrder{
			value: "desc",
		},
	}
}

func (c ShowAssetMetaRequestOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowAssetMetaRequestOrder) UnmarshalJSON(b []byte) error {
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
