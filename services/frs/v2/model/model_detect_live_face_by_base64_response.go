package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectLiveFaceByBase64Response struct {
	Result *LiveDetectFaceRespResult `json:"result,omitempty"`
	// 警告信息列表。 调用失败时无此字段。

	WarningList    *[]WarningList `json:"warning-list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DetectLiveFaceByBase64Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByBase64Response struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByBase64Response", string(data)}, " ")
}
