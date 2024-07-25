package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageBatchSyncDetectionResult struct {

	// 审核结果是否通过。 - block：包含敏感信息，不通过。 - pass：不包含敏感信息，通过 。 - review：需要人工复检。
	Suggestion *string `json:"suggestion,omitempty"`

	// 检测结果的一级标签。支持category列表如下： - terrorism: 暴恐 - porn: 色情 - image_text: 图文审核
	Category *string `json:"category,omitempty"`

	// 检测详情。
	Details *[]ImageDetectionResultDetail `json:"details,omitempty"`

	// 图文审核检测出的文本，只有在category参数配置image_text且检测出文本时展示该字段。
	OcrText *string `json:"ocr_text,omitempty"`

	// 图片唯一标识。同一次请求中不可重复，由大小写英文字母、数字、下划线（_）、中划线（-）组成，不超过30个字符。
	DataId string `json:"data_id"`

	// 调用失败时的错误码，具体请参见错误码。 调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。 调用成功时无此字段。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ImageBatchSyncDetectionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBatchSyncDetectionResult struct{}"
	}

	return strings.Join([]string{"ImageBatchSyncDetectionResult", string(data)}, " ")
}
