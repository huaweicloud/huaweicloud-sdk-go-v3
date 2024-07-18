package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimPersonalTemplateContentAction 个人模板事件对象。
type AimPersonalTemplateContentAction struct {

	// 此字段根据action_type对应不同的含义，具体对应如下。  - action_type=OPEN_URL：表示H5访问地址。必须为HTTPS，支持含动态参数。字符长度为1-1000。示例：https://XXXXX/${param1} - action_type=OPEN_QUICK：表示快应用deeplink地址。支持含动态参数，字符长度为1-1000。示例：hap://app/xxx/${param1} - action_type=OPEN_APP：表示APP的deeplink地址。支持含动态参数，字符长度为1-1000。示例：weixin:// - action_type=DIAL_PHONE：表示电话号码。不能超过20个字符。示例：18600000000 - action_type=OPEN_SMS：表示电话号码。不能超过20个字符。示例：18600000000 - action_type=OPEN_EMAIL：表示邮箱地址。不能超过100个字符。示例：1046520406@qq.com - action_type=OPEN_SCHEDULE：表示日程标题。不能超过100个字符。示例：日常需求评审 - action_type=OPEN_MAP：表示位置名。不能超过100个字符。示例：龙泰利科技大厦 - action_type=OPEN_BROWSER：表示网址。支持HTTPS或HTTP，支持含动态参数，不能超过1000个字符。示例：https://XXXXX/${param1} - action_type=OPEN_POPUP：表示弹窗标题。不能超过30个字符。参数示例：xxx商品 - action_type=COPY_PARAMETER：表示复制的内容。支持含动态参数，不能超过20个字符。复制验证码示例：83721 - action_type=VIEW_PIC：表示要打开的大图ID。配置在打开大图的资源地址与模板上的图片资源地址一致。如果模板资源类型是ID，则传ID，如果是资源地址，则使用资源地址。最大长度不能超过1000个字符。例如：当src_type为1时，传入ID：691996319597764608。当src_type为2时，使用资源地址：https://www.xxxx.cn/src/image/head.jpg
	Target string `json:"target"`

	// 弹窗内容。  > action_type=OPEN_POPUP为必填。不能超过100个字符。示例：是否喜欢该商品。
	Content *string `json:"content,omitempty"`

	// 包名。  > action_type=OPEN_APP为必填。不能超过50个字符。示例：com.xxxx.service.koomsg。
	PackageName *string `json:"package_name,omitempty"`

	// 兜底URL。支持快应用deeplink或H5的HTTPS网址，不能超过1000个字。  > - action_type=OPEN_APP为选填，其他类型不填 > - 兜底类型为0时，可不填 > - 当兜底类型为2并且提交厂商列表中包含OPPO厂商时为必填
	FloorUrl *string `json:"floor_url,omitempty"`

	// 兜底类型。如果传入的厂商不支持该兜底类型，接口会返回错误。如果不传入厂商，则不对兜底类型进行校。 - 0：打开应用市场 - 1：打开H5页面（通过收件箱内置浏览器打开） - 2：打开浏览器 - 3：打开快应用  > action_type=OPEN_APP为选填，其他类型不填；action_type=OPEN_APP时此参数不填则默认打开应用市场。打开链接为http格式时必须选择打开浏览器；打开链接为https格式且内容只是一个普通页面时，可以使用打开H5页面，当链接中有下载指引或打开小程序由于部分内置浏览器功能不全可能导致打开异常，建议使用打开浏览器，请按需选择兜底类型。 > - 华为：支持以上4种兜底 > - 魅族：支持以上4种兜底 > - 小米：不支持打开H5页面兜底 > - OPPO：不支持打开快应用兜底 > - VIVO：不支持打开快应用兜底 > - 三星：不支持打开应用市场和打开浏览器。当创建的模板仅包含三星厂商时，兜底URL不支持打开浏览器和打开应用市场；当创建的模板包含三星和其它厂商时，以其它厂商的限制为准，三星的兜底链接将不生效
	FloorType *int32 `json:"floor_type,omitempty"`

	// 邮件标题。  > action_type=OPEN_EMAIL为必填。不能超过100个字符。示例：618活动促销。
	Subject *string `json:"subject,omitempty"`

	// 邮件正文/短信正文。  > action_type=OPEN_SMS或OPEN_EMAIL为必填。不能超过100个字符。 > > 短信正文示例1：今天回家吃饭吗； > > 邮件正文示例2：您有一张优惠券领取。
	Body *string `json:"body,omitempty"`

	// 日程内容描述。  > action_type=OPEN_SCHEDULE为必填。不能超过100个字符。示例：评审这个月版本需求。
	Description *string `json:"description,omitempty"`

	// 日程开始时间。格式为：yyyy-MM-dd HH:mm:ss。  > 当action_type=OPEN_SCHEDULE时为必填。
	BeginTime *string `json:"begin_time,omitempty"`

	// 日程结束时间。格式为：yyyy-MM-dd HH:mm:ss。  > 当action_type=OPEN_SCHEDULE时为必填。
	EndTime *string `json:"end_time,omitempty"`

	// 地址的详细说明。  > action_type=OPEN_MAP为必填。不能超过100个字符。示例：高新中四道龙泰利科技大厦。
	Address *string `json:"address,omitempty"`

	// 地图经度。  > action_type=OPEN_MAP为必填。不能超过20个字符。示例：113.941618。
	Longitude *string `json:"longitude,omitempty"`

	// 地图纬度。  > action_type=OPEN_MAP为必填。不能超过20个字符。示例：22.548804。
	Latitude *string `json:"latitude,omitempty"`

	// 按钮展示文本。  > action_type=OPEN_POPUP为必填。不能超过12个字符。示例：确定。
	TextButton *string `json:"text_button,omitempty"`

	// 弹窗模态。  - 0：模态（默认） - 1：非模态（暂不支持）  > action_type=OPEN_POPUP为必填。
	Mode *int32 `json:"mode,omitempty"`
}

func (o AimPersonalTemplateContentAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimPersonalTemplateContentAction struct{}"
	}

	return strings.Join([]string{"AimPersonalTemplateContentAction", string(data)}, " ")
}
