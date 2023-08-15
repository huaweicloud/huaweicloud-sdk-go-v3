package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectLiveByUrlIntlResponse Response Object
type DetectLiveByUrlIntlResponse struct {
	VideoResult *LiveDetectRespVideoresult `json:"video-result,omitempty"`

	// [警告信息列表，WarningList结构见[WarningList](https://support.huaweicloud.com/api-face/face_02_0077.html)。调用失败时无此字段](tag:hc) [警告信息列表，WarningList结构见[WarningList](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0077.html)。调用失败时无此字段](tag:hk)
	WarningList    *[]WarningList `json:"warning-list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DetectLiveByUrlIntlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByUrlIntlResponse struct{}"
	}

	return strings.Join([]string{"DetectLiveByUrlIntlResponse", string(data)}, " ")
}
