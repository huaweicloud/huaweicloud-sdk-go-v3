package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 命名实体类
type DomainNamedEntity struct {

	// 实体文本。
	Word string `json:"word"`

	// 实体类型，枚举类型。 通用领域：支持人名nr，地名ns，机构名nt，时间点tpt，日期day，百分比pct，货币额度mny，序数词ord，计量规格词qtt，民族race，职业job，邮箱email，国家coun，节日fest。 商务领域：支持公司名com、品牌名bra、职业job、职位post、邮箱email、手机号码cell、电话号码tele、IP地址ip、身份证号id、网址web。 娱乐领域：支持电影名mov、动漫anime、书名book、互联网int、歌名song、产品名pro、电视剧名dra、电视节目名tv。
	Tag string `json:"tag"`

	// 实体文本在待分析文本中的起始位置。
	Offset int32 `json:"offset"`

	// 实体文本长度。
	Len int32 `json:"len"`
}

func (o DomainNamedEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNamedEntity struct{}"
	}

	return strings.Join([]string{"DomainNamedEntity", string(data)}, " ")
}
