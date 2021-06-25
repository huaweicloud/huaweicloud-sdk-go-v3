package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 媒资状态.<br/>
type Status struct {
	value string
}

type StatusEnum struct {
	UNCREATED          Status
	DELETED            Status
	CANCELLED          Status
	SERVER_ERROR       Status
	UPLOAD_FAILED      Status
	CREATING           Status
	PUBLISHED          Status
	WAITING_TRANSCODE  Status
	TRANSCODING        Status
	TRANSCODE_FAILED   Status
	TRANSCODE_SUCCEED  Status
	CREATED            Status
	UNPUBLISHED        Status
	NO_ASSET           Status
	DELETING           Status
	DELETE_FAILED      Status
	OBS_CREATING       Status
	OBS_CREATE_FAILED  Status
	OBS_CREATE_SUCCESS Status
	U_NK_NOW           Status
}

func GetStatusEnum() StatusEnum {
	return StatusEnum{
		UNCREATED: Status{
			value: "UNCREATED",
		},
		DELETED: Status{
			value: "DELETED",
		},
		CANCELLED: Status{
			value: "CANCELLED",
		},
		SERVER_ERROR: Status{
			value: "SERVER_ERROR",
		},
		UPLOAD_FAILED: Status{
			value: "UPLOAD_FAILED",
		},
		CREATING: Status{
			value: "CREATING",
		},
		PUBLISHED: Status{
			value: "PUBLISHED",
		},
		WAITING_TRANSCODE: Status{
			value: "WAITING_TRANSCODE",
		},
		TRANSCODING: Status{
			value: "TRANSCODING",
		},
		TRANSCODE_FAILED: Status{
			value: "TRANSCODE_FAILED",
		},
		TRANSCODE_SUCCEED: Status{
			value: "TRANSCODE_SUCCEED",
		},
		CREATED: Status{
			value: "CREATED",
		},
		UNPUBLISHED: Status{
			value: "UNPUBLISHED",
		},
		NO_ASSET: Status{
			value: "NO_ASSET",
		},
		DELETING: Status{
			value: "DELETING",
		},
		DELETE_FAILED: Status{
			value: "DELETE_FAILED",
		},
		OBS_CREATING: Status{
			value: "OBS_CREATING",
		},
		OBS_CREATE_FAILED: Status{
			value: "OBS_CREATE_FAILED",
		},
		OBS_CREATE_SUCCESS: Status{
			value: "OBS_CREATE_SUCCESS",
		},
		U_NK_NOW: Status{
			value: "UNkNOW",
		},
	}
}

func (c Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *Status) UnmarshalJSON(b []byte) error {
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
