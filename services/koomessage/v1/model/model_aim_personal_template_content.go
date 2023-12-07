package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimPersonalTemplateContent 模板协议内容数据结构。
type AimPersonalTemplateContent struct {

	// 模板资源类型。 - text：表示文本 - image：表示图片 - video：表示视频 - button：表示按钮 - followPub：表示华为服务号，暂不支持  > 图片轮播类模板最多可以放5张图片，即card_id为CarouselSquareImage、CarouselImageSixteenToNine、CarouselVerticalImage时，type为image的资源最多有5个。
	Type string `json:"type"`

	// 资源类型为Text或Button时，为必填。文本长度限制请按智能短信模板版式格式标准。  > 智能信息模板标准版式要求可登录KooMessage控制台，在创建智能信息模板中查看。
	Content *string `json:"content,omitempty"`

	// 子内容。非必填，文本长度限制请参考智能信息模板标准版式要求。  > 智能信息模板标准版式要求可登录KooMessage控制台，在创建智能信息模板中查看。
	ContentChild *string `json:"content_child,omitempty"`

	// src类型。资源类型为Image或Video时，该项为必填项。 - 1：指资源ID  - 2：指资源地址  > src_type为2，即资源地址时，src或cover内容必须是存储在客户侧服务器上的资源地址。
	SrcType *int32 `json:"src_type,omitempty"`

	// 资源来源，资源类型为Image或Video时，为必填，通过设置视频模板封面图接口设置视频封面。 > - 如上src_type为1，即资源ID时，参数填入上传模板素材接口中返回的aim_resource_id，如：691996319597764608 > - 如上src_type为2，即资源地址时，参数填写资源完整的URL，最大长度不超过1000个字符
	Src *string `json:"src,omitempty"`

	// 视频封面。 > 资源类型为Video时，为必填。 > - 如上src_type为1，即资源ID时，参数填入上传模板素材接口中返回的aim_resource_id，如：691996319597764608 > - 如上src_type为2，即资源地址时，参数填写资源完整的URL，最大长度不超过1000个字符
	Cover *string `json:"cover,omitempty"`

	// 是否为文本标题。  - true：是 - false：不是  > 不填默认为false。
	IsTextTitle *string `json:"is_text_title,omitempty"`

	// 功能类型。  - OPEN_URL：表示跳转H5 - OPEN_QUICK：表示跳转快应用 - OPEN_APP：表示跳转APP - DIAL_PHONE：表示拉起拨号盘 - OPEN_SMS：表示新建短信息 - OPEN_EMAIL：表示打开邮箱 - OPEN_SCHEDULE：表示新建日程 - OPEN_MAP：表示位置定位 - OPEN_BROWSER：表示打开浏览器 - OPEN_POPUP：表示弹窗 - COPY_PARAMETER：表示复制 - VIEW_PIC：表示打开大图  > - type为Image和Button时为必填项，必须绑定事件 > - type为其他类型时则不必填 > - OPPO厂商点击事件类型只支持打开浏览器、打开快应用、打开APP、跳转H5 > - VIVO厂商点击事件类型不支持打开邮箱、打开地图 > - MEIZU厂商点击事件类型不支持打开大图 > - 横滑类1、横滑类2版式的图片不支持绑定事件，默认与按钮事件一致 > - 三星厂商点击事件类型不支持新建日程、打开大图
	ActionType *string `json:"action_type,omitempty"`

	// 卡片组件的位置序号。 > 资源在卡片上相对的位置序号，按照优先从左到右，再从上到下的编排原则，统一编号。
	PositionNumber int32 `json:"position_number"`

	// 组件是否可见。 - 0：隐藏（某些组件可设置隐藏） - 1：可见  > 目前仅针对电商多商品（Ecommerce）、多卡券（CardVouchers）、增强机票类（PlaneTrip）这三种版式起效。
	Visible *int32 `json:"visible,omitempty"`

	// 是否显示货币符号。  - 0：隐藏  - 1：可见  > 当模板为电商类时是否显示￥符号，默认可见。
	CurrencyDisplay *int32 `json:"currency_display,omitempty"`

	// OPPO红包背景。  > - 当src_type为1时，即资源ID时，参数填入上传模板素材接口中返回的aim_resource_id，如：691996319597764608 > - 当src_type为2时，即资源地址时，参数填写资源完整的URL，最大长度不超过1000个字符 >- 当模板为红包类，即card_id为RedPacket时用于指定OPPO厂商红包背景图，具体使用可参考创建红包类模板请求示例
	OppoBackground *string `json:"oppo_background,omitempty"`

	// VIVO红包背景。  > - 当src_type为1时，即资源ID时，参数填入上传模板素材接口中返回的aim_resource_id，如：691996319597764608 > - 当src_type为2时，即资源地址时，参数填写资源完整的URL，最大长度不超过1000个字符
	VivoBackground *string `json:"vivo_background,omitempty"`

	// 表示短视频模板视频和封面的宽高比。即card_id为ShortVideo时，此项有值。 - threeToFour: 宽高比为3:4 - oneToOne: 宽高比为1:1
	Ratio *string `json:"ratio,omitempty"`

	Action *AimPersonalTemplateContentAction `json:"action,omitempty"`

	// 当模板为电商领券类竖版，即card_id为EcommerceCouponVertical时用于指定按钮类型，具体使用可参考创建电商领券类竖版模板请求示例。 - static：静态按钮 - dynamic：动态按钮
	ButtonType *string `json:"button_type,omitempty"`
}

func (o AimPersonalTemplateContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimPersonalTemplateContent struct{}"
	}

	return strings.Join([]string{"AimPersonalTemplateContent", string(data)}, " ")
}
