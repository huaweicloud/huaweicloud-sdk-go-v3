package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectLiveFaceByUrlResponse struct {
	Result *LiveDetectFaceRespResult `json:"result,omitempty" xml:"result"`

	// 警告信息列表。 调用失败时无此字段。
	WarningList    *[]WarningList `json:"warning-list,omitempty" xml:"warning-list"`
	HttpStatusCode int            `json:"-"`
}

func (o DetectLiveFaceByUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByUrlResponse struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByUrlResponse", string(data)}, " ")
}
