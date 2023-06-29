package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunGetFileTranslationResultResponse Response Object
type RunGetFileTranslationResultResponse struct {

	// 当前翻译状态。具体状态如下所示： WAITING 等待翻译。 FINISHED 翻译已经完成。 ERROR 翻译过程中发生错误。 调用失败时无此字段。
	Status *string `json:"status,omitempty"`

	// 临时url，用于获取翻译结果，有效期十分钟。过期后请再次调用接口获取新的url。调用失败时或翻译状态非FINISHED时无此字段。
	Url *string `json:"url,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunGetFileTranslationResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunGetFileTranslationResultResponse struct{}"
	}

	return strings.Join([]string{"RunGetFileTranslationResultResponse", string(data)}, " ")
}
