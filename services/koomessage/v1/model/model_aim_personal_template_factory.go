package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimPersonalTemplateFactory 厂商。
type AimPersonalTemplateFactory struct {

	// 厂商类型。  - HUAWEI：华为 - XIAOMI：小米 - OPPO：OPPO - VIVO：VIVO - MEIZU：魅族 - SAMSUNG：三星  > 各厂商支持的布局类型，布局类型与card_id字段相对应。分别如下： > - HUAWEI：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、长文本类（PureText）、横滑类1（Carousel）、横滑类2（CarouselTitle）、视频图文类（VideoImageAndText）、视频类（Video）、电商类（ECImageAndText）、红包类（RedPacket）、个性化红包类（RedPacketPersonal）、增强通知类（Notification2）、图片轮播类1:1（CarouselSquareImage）、图片轮播类16:9（CarouselImageSixteenToNine）、图片轮播类48:65（CarouselVerticalImage）、图文视频类（ImageTextAndVideo）、一般通知类（Notification1）、单卡券（CardVoucher）、多卡券（CardVouchers）、电商多商品类（Ecommerce）、电商领券类竖版（EcommerceCouponVertical）、机票类（Trip1）、火车票类（Trip2）、汽车票类（Trip3）、增强机票类（PlaneTrip）、海报类（SimplePoster）、超文本普通类（NativePureText）、超文本增强类（NativeImageAndText）、短剧视频类（ShortVideo） > - XIAOMI：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、红包类（RedPacket）、增强通知类（Notification2）、一般通知类（Notification1） > - OPPO：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、长文本类（PureText）、横滑类（Carousel）、视频类（Video）、电商类（ECImageAndText）、红包类（RedPacket）、图片轮播类1:1（CarouselSquareImage）、图片轮播类16:9（CarouselImageSixteenToNine）、图片轮播类48:65（CarouselVerticalImage）、单卡券（CardVoucher） > - MEIZU：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、横滑类1（Carousel）、横滑类2（CarouselTitle） > - VIVO：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、图片轮播类1:1（CarouselSquareImage）、图片轮播类16:9（CarouselImageSixteenToNine）、图片轮播类48:65（CarouselVerticalImage）、视频类（Video）、电商类（ECImageAndText）、红包类（RedPacket）、增强通知类（Notification2）、一般通知类（Notification1）、个性化红包类（RedPacketPersonal）、单卡券（CardVoucher） > - 三星：多图文类（MultipleImageAndText）、图文类（StandardImageAndText）、长文本类（PureText）、横滑类1（Carousel）、横滑类2（CarouselTitle）、视频图文类（VideoImageAndText）、视频类（Video）、电商类（ECImageAndText）、红包类（RedPacket）、图片轮播类1:1（CarouselSquareImage）、图片轮播类16:9（CarouselImageSixteenToNine）、图片轮播类48:65（CarouselVerticalImage）、图文视频类（ImageTextAndVideo）、一般通知类（Notification1）、机票类（Trip1）、火车票类（Trip2）、汽车票类（Trip3）
	FactoryType string `json:"factory_type"`

	// 支持状态。 - 1：支持 - 0：不支持
	State int32 `json:"state"`
}

func (o AimPersonalTemplateFactory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimPersonalTemplateFactory struct{}"
	}

	return strings.Join([]string{"AimPersonalTemplateFactory", string(data)}, " ")
}
