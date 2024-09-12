package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveEvent struct {

	// **参数解释**： 事件戳。从1970-01-01 00:00:00:000开始的毫秒数
	Timestamp int64 `json:"timestamp"`

	// **参数解释**： 事件类型。 * 1 弹幕信息 * 2 用户入场 * 3 用户点赞 * 4 用户送礼 * 5 用户订阅\\关注 * 6 房间观看人数
	Type *int32 `json:"type,omitempty"`

	// 事件内容。 事件类型不同，content内容也不同，定义如下所示： - type=1 弹幕信息   content定义：     * user:用户，json     * content: string,弹幕内容     user定义：     * userId：用户id，string     * name：用户姓名，string     * level：用户平台等级，int     * badge：用户牌子名称，string     * badgeLevel：牌子等级，int     举例：   ```json   {     \"timestamp\": 1694481224245,     \"type\": 1,     \"content\": \"{\\\"user\\\":{\\\"userId\\\":\\\"2027271526\\\",\\\"name\\\":\\\"***\\\",\\\"level\\\":17,\\\"badge\\\":\\\"\\\",\\\"badgeLevel\\\":0},\\\"content\\\":\\\"***\\\"}\"   }   ``` - type=2 用户入场   content定义：     * user:用户，json     举例：   ```json   {     \"timestamp\": 1694481227655,     \"type\": 2,     \"content\": \"{\\\"user\\\":{\\\"userId\\\":\\\"2978899271\\\",\\\"name\\\":\\\"***\\\",\\\"level\\\":1,\\\"badge\\\":\\\"\\\",\\\"badgeLevel\\\":0}}\"   }   ``` - type=3 用户点赞    content定义：     * user：用户，json     举例：   ```json   {     \"timestamp\": 1694481227655,     \"type\": 2,     \"content\": \"{\\\"user\\\":{\\\"userId\\\":\\\"2978899271\\\",\\\"name\\\":\\\"***\\\",\\\"level\\\":1,\\\"badge\\\":\\\"\\\",\\\"badgeLevel\\\":0}}\"   }   ``` - type=4 用户送礼   content定义：     * user：用户，json     * gift：礼物信息,json     gift定义：     * giftName：礼物名称，string     * giftNum：礼物数量，int     * totalValue：此处礼物总价值，int     * giftValue：单个礼物价值，int     举例：   ```json   {     \"timestamp\": 1690531652862,     \"type\": 4,     \"content\": \"{\\\"gift\\\":{\\\"giftName\\\":\\\"小星星\\\",\\\"giftNum\\\":10,\\\"totalValue\\\":10,\\\"giftValue\\\":3},\\\"user\\\":{\\\"userId\\\":\\\"douyin_95882940927\\\",\\\"name\\\":\\\"纯爱战士熙熙\\\",\\\"level\\\":2,\\\"badge\\\":\\\"\\\",\\\"badgeLevel\\\":0}}\"   }   ``` - type=5 用户订阅    暂未使用 - type=6 房间观看人数   content定义：     * numberOfViewers：房间观看人数，int     举例：   ```json   {     \"timestamp\": 1694481236886,     \"type\": 6,     \"content\": \"{\\\"numberOfViewers\\\":5466}\"   }   ```
	Content *string `json:"content,omitempty"`
}

func (o LiveEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveEvent struct{}"
	}

	return strings.Join([]string{"LiveEvent", string(data)}, " ")
}
