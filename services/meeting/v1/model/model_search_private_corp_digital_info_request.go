package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPrivateCorpDigitalInfoRequest Request Object
type SearchPrivateCorpDigitalInfoRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	// 被绑定的登录账号或realNameAccount经过base64编码后的信息。
	Account string `json:"account"`

	// 传译语言。
	Language string `json:"language"`
}

func (o SearchPrivateCorpDigitalInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPrivateCorpDigitalInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchPrivateCorpDigitalInfoRequest", string(data)}, " ")
}
