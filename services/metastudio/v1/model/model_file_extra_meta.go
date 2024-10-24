package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FileExtraMeta 文件数据。
type FileExtraMeta struct {

	// 视频转码状态。 * WAITING：等待 * TRANSCODING：转码中 * FAILED：失败 * SUCCEEDED：成功
	VideoTranscodingStatus *FileExtraMetaVideoTranscodingStatus `json:"video_transcoding_status,omitempty"`
}

func (o FileExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileExtraMeta struct{}"
	}

	return strings.Join([]string{"FileExtraMeta", string(data)}, " ")
}

type FileExtraMetaVideoTranscodingStatus struct {
	value string
}

type FileExtraMetaVideoTranscodingStatusEnum struct {
	WAITING     FileExtraMetaVideoTranscodingStatus
	TRANSCODING FileExtraMetaVideoTranscodingStatus
	FAILED      FileExtraMetaVideoTranscodingStatus
	SUCCEEDED   FileExtraMetaVideoTranscodingStatus
}

func GetFileExtraMetaVideoTranscodingStatusEnum() FileExtraMetaVideoTranscodingStatusEnum {
	return FileExtraMetaVideoTranscodingStatusEnum{
		WAITING: FileExtraMetaVideoTranscodingStatus{
			value: "WAITING",
		},
		TRANSCODING: FileExtraMetaVideoTranscodingStatus{
			value: "TRANSCODING",
		},
		FAILED: FileExtraMetaVideoTranscodingStatus{
			value: "FAILED",
		},
		SUCCEEDED: FileExtraMetaVideoTranscodingStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c FileExtraMetaVideoTranscodingStatus) Value() string {
	return c.value
}

func (c FileExtraMetaVideoTranscodingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileExtraMetaVideoTranscodingStatus) UnmarshalJSON(b []byte) error {
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
