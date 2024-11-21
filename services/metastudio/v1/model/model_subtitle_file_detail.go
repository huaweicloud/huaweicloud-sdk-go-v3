package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubtitleFileDetail struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	// 字幕文件生成状态。 * GENERATING：字幕文件生成中。 * GENERATE_SUCCEED：字幕文件生成成功。 * GENERATE_FAILED：字幕文件生成失败。
	SubtitleFileState *SubtitleFileDetailSubtitleFileState `json:"subtitle_file_state,omitempty"`

	// 字幕文件下载链接。
	SubtitleFileDownloadUrl *string `json:"subtitle_file_download_url,omitempty"`

	// 字幕文件上传链接。
	SubtitleFileUploadUrl *string `json:"subtitle_file_upload_url,omitempty"`

	// 字幕文件生成时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	GenerateTime *string `json:"generate_time,omitempty"`
}

func (o SubtitleFileDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleFileDetail struct{}"
	}

	return strings.Join([]string{"SubtitleFileDetail", string(data)}, " ")
}

type SubtitleFileDetailSubtitleFileState struct {
	value string
}

type SubtitleFileDetailSubtitleFileStateEnum struct {
	GENERATE_SUCCEED SubtitleFileDetailSubtitleFileState
	GENERATE_FAILED  SubtitleFileDetailSubtitleFileState
	GENERATING       SubtitleFileDetailSubtitleFileState
}

func GetSubtitleFileDetailSubtitleFileStateEnum() SubtitleFileDetailSubtitleFileStateEnum {
	return SubtitleFileDetailSubtitleFileStateEnum{
		GENERATE_SUCCEED: SubtitleFileDetailSubtitleFileState{
			value: "GENERATE_SUCCEED",
		},
		GENERATE_FAILED: SubtitleFileDetailSubtitleFileState{
			value: "GENERATE_FAILED",
		},
		GENERATING: SubtitleFileDetailSubtitleFileState{
			value: "GENERATING",
		},
	}
}

func (c SubtitleFileDetailSubtitleFileState) Value() string {
	return c.value
}

func (c SubtitleFileDetailSubtitleFileState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubtitleFileDetailSubtitleFileState) UnmarshalJSON(b []byte) error {
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
