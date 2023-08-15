package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockPortRequestBody 通道号绑定服务号请求体。
type LockPortRequestBody struct {

	// 服务号ID。
	PubId string `json:"pub_id"`

	// 通道号。 - port_type=5时，长度必须为5 - port_type=1或3，长度在21位内
	Port string `json:"port"`

	// 通道号绑定/解绑的province字段取值范围如下： 全国、河北省、山西省、辽宁省、吉林省、黑龙江省、江苏省、浙江省、安徽省、福建省、江西省、山东省、河南省、湖北省、湖南省、广东省、海南省、四川省、贵州省、云南省、陕西省、甘肃省、青海省、台湾省、内蒙古自治区、广西壮族自治区、西藏自治区、宁夏回族自治区、新疆维吾尔自治区、北京市、天津市、上海市、重庆市、香港特别行政区、澳门特别行政区。
	Province string `json:"province"`

	// 绑定签名，必须是该端口号签名的子集。单个签名长度为2-18。
	Sign []string `json:"sign"`

	// 关联通道号ID，取通道号列表返回的ID。
	ExtPortId string `json:"ext_port_id"`
}

func (o LockPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockPortRequestBody struct{}"
	}

	return strings.Join([]string{"LockPortRequestBody", string(data)}, " ")
}
