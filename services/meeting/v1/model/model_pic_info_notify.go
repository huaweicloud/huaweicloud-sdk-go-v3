package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PicInfoNotify 子画面信息。
type PicInfoNotify struct {

	// 多画面中每个画面的编号，编号从1开始。
	Index *int32 `json:"index,omitempty"`

	// 每个画面中的与会者SIP号码。SIP号码可以通过[[查询企业通讯](https://support.huaweicloud.com/api-meeting/meeting_21_0512.html)](tag:hws)[[查询企业通讯](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0512.html)](tag:hk)获取。
	Id *[]string `json:"id,omitempty"`

	// 是否为辅流。 * 0： 不是辅流 * 1： 是辅流
	Share *int32 `json:"share,omitempty"`
}

func (o PicInfoNotify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicInfoNotify struct{}"
	}

	return strings.Join([]string{"PicInfoNotify", string(data)}, " ")
}
