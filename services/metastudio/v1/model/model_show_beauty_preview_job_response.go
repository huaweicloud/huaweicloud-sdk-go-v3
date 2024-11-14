package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBeautyPreviewJobResponse Response Object
type ShowBeautyPreviewJobResponse struct {

	// 任务的状态。 * WAIT_IMAGE_UPLOAD：待上传美白前图片 * WAITING：等待生成美白预览图 * PROCESSING：美白预览图生成中 * SUCCESS：美白预览图生成成功 * FAILED：美白预览图生成失败
	State *ShowBeautyPreviewJobResponseState `json:"state,omitempty"`

	// 美白后图片下载url。
	PostBeautyImageDownloadUrl *string `json:"post_beauty_image_download_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBeautyPreviewJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBeautyPreviewJobResponse struct{}"
	}

	return strings.Join([]string{"ShowBeautyPreviewJobResponse", string(data)}, " ")
}

type ShowBeautyPreviewJobResponseState struct {
	value string
}

type ShowBeautyPreviewJobResponseStateEnum struct {
	WAIT_IMAGE_UPLOAD ShowBeautyPreviewJobResponseState
	WAITING           ShowBeautyPreviewJobResponseState
	PROCESSING        ShowBeautyPreviewJobResponseState
	SUCCESS           ShowBeautyPreviewJobResponseState
	FAILED            ShowBeautyPreviewJobResponseState
}

func GetShowBeautyPreviewJobResponseStateEnum() ShowBeautyPreviewJobResponseStateEnum {
	return ShowBeautyPreviewJobResponseStateEnum{
		WAIT_IMAGE_UPLOAD: ShowBeautyPreviewJobResponseState{
			value: "WAIT_IMAGE_UPLOAD",
		},
		WAITING: ShowBeautyPreviewJobResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowBeautyPreviewJobResponseState{
			value: "PROCESSING",
		},
		SUCCESS: ShowBeautyPreviewJobResponseState{
			value: "SUCCESS",
		},
		FAILED: ShowBeautyPreviewJobResponseState{
			value: "FAILED",
		},
	}
}

func (c ShowBeautyPreviewJobResponseState) Value() string {
	return c.value
}

func (c ShowBeautyPreviewJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBeautyPreviewJobResponseState) UnmarshalJSON(b []byte) error {
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
