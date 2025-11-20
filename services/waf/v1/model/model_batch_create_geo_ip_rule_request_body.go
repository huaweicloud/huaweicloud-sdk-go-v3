package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateGeoIpRuleRequestBody struct {

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 地理位置封禁区域： (CA： 加拿大,US： 美国,AU： 澳大利亚,IN： 印度,JP： 日本,UK： 英国,FR： 法国,DE： 德国,BR： 巴西,Thailand： 泰国,Singapore： 新加坡,South Africa： 南非,Mexico： 墨西哥,Peru： 秘鲁,Indonesia： 印度尼西亚,GD： 广东,FJ： 福建,JL： 吉林,LN： 辽宁,TW： 中国-台湾,GZ： 贵州,AH： 安徽,HL： 黑龙江,HA： 河南,SC： 四川,HE： 河北,YN： 云南,HB： 湖北,HI： 海南,QH： 青海,HN： 湖南,JX： 江西,SX： 山西,SN： 陕西,ZJ： 浙江,GS： 甘肃,JS： 江苏,SD： 山东,BJ： 北京,SH： 上海,TJ： 天津,CQ： 重庆,MO： 中国-澳门,HK： 中国-香港,NX： 宁夏,GX： 广西,XJ： 新疆,XZ： 西藏,NM： 内蒙古)。具体的位置编码参见《附录-地理位置编码》查询。
	Geoip string `json:"geoip"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White int32 `json:"white"`

	// ip范围。若您的网站使用独享模式，请确认独享引擎是否全部升级到最新版本，避免造成异常。202412之后的版本支持配置IP范围
	IpType string `json:"ip_type"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 需要添加规则的策略id
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateGeoIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGeoIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateGeoIpRuleRequestBody", string(data)}, " ")
}
