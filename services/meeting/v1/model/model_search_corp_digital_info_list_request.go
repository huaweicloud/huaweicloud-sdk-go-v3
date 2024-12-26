package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpDigitalInfoListRequest Request Object
type SearchCorpDigitalInfoListRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	// 查询类型，PUBLIC：公共、LOCAL：本地会议。
	Type string `json:"type"`

	// 传译语言。
	Language string `json:"language"`
}

func (o SearchCorpDigitalInfoListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDigitalInfoListRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpDigitalInfoListRequest", string(data)}, " ")
}
