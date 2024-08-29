package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubtitleFileInfo struct {

	// 字幕文件下载链接。
	SubtitleFileDownloadUrl *string `json:"subtitle_file_download_url,omitempty"`

	// 字幕文件上传链接。
	SubtitleFileUploadUrl *string `json:"subtitle_file_upload_url,omitempty"`

	// 字幕文件生成状态。 * GENERATING：字幕文件生成中。 * GENERATE_SUCCEED：字幕文件生成成功。 * GENERATE_FAILED：字幕文件生成失败。
	SubtitleFileState *SubtitleFileInfoSubtitleFileState `json:"subtitle_file_state,omitempty"`

	// 字幕文件生成任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o SubtitleFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleFileInfo struct{}"
	}

	return strings.Join([]string{"SubtitleFileInfo", string(data)}, " ")
}

type SubtitleFileInfoSubtitleFileState struct {
	value string
}

type SubtitleFileInfoSubtitleFileStateEnum struct {
	GENERATE_SUCCEED SubtitleFileInfoSubtitleFileState
	GENERATE_FAILED  SubtitleFileInfoSubtitleFileState
	GENERATING       SubtitleFileInfoSubtitleFileState
}

func GetSubtitleFileInfoSubtitleFileStateEnum() SubtitleFileInfoSubtitleFileStateEnum {
	return SubtitleFileInfoSubtitleFileStateEnum{
		GENERATE_SUCCEED: SubtitleFileInfoSubtitleFileState{
			value: "GENERATE_SUCCEED",
		},
		GENERATE_FAILED: SubtitleFileInfoSubtitleFileState{
			value: "GENERATE_FAILED",
		},
		GENERATING: SubtitleFileInfoSubtitleFileState{
			value: "GENERATING",
		},
	}
}

func (c SubtitleFileInfoSubtitleFileState) Value() string {
	return c.value
}

func (c SubtitleFileInfoSubtitleFileState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubtitleFileInfoSubtitleFileState) UnmarshalJSON(b []byte) error {
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
