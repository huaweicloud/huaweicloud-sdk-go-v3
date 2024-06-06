package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchOptionalParam 搜索的可选参数，其中每个参数仅对部分服务类型生效，具体可参见服务类型说明。
type SearchOptionalParam struct {

	// 是否进行目标检测，默认为true。
	DoDet *bool `json:"do_det,omitempty"`

	// 目标矩形框坐标，如给定则不进行目标检测，直接使用该box作为目标。格式为“x1,y1,x2,y2”（无空格），x1/y1为目标左上角坐标，x2/y2为目标右下角坐标，具体要求如下： - 0 <= x1 < x2 <= width，默认要求x2-x1 >= 15，具体可参考服务类型说明。 - 0 <= y1 < y2 <= height，默认要求y2-y1 >= 15，具体可参考服务类型说明。
	Box *string `json:"box,omitempty"`

	// 是否进行对象分类，默认为true。
	DoCls *bool `json:"do_cls,omitempty"`

	// 对象类目，如给定则不进行对象分类，直接使用该category作为类目。具体类目信息可参见对应的服务类型说明。
	Category *int32 `json:"category,omitempty"`

	// 去重标签名，必须为服务实例custom_tags中已存在的key。 - 如给定则会对该key下相同value的数据进行去重，仅保留得分最高的数据。 - 针对没有设置该标签的数据，会直接过滤。
	CollapseKey *string `json:"collapse_key,omitempty"`

	// 扫描节点上限。值越大精度越高，查询速度变慢。默认值为10000。
	MaxScanNum *int32 `json:"max_scan_num,omitempty"`

	// 查询考察中心点的数目。值越大精度越高，查询速度变慢。默认值为100。
	Nprobe *int32 `json:"nprobe,omitempty"`

	// 文本字符串的语言类型枚举值。
	TextLang *SearchOptionalParamTextLang `json:"text_lang,omitempty"`
}

func (o SearchOptionalParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchOptionalParam struct{}"
	}

	return strings.Join([]string{"SearchOptionalParam", string(data)}, " ")
}

type SearchOptionalParamTextLang struct {
	value string
}

type SearchOptionalParamTextLangEnum struct {
	AR    SearchOptionalParamTextLang
	DE    SearchOptionalParamTextLang
	RU    SearchOptionalParamTextLang
	FR    SearchOptionalParamTextLang
	KO    SearchOptionalParamTextLang
	PT    SearchOptionalParamTextLang
	JA    SearchOptionalParamTextLang
	TH    SearchOptionalParamTextLang
	TR    SearchOptionalParamTextLang
	ES    SearchOptionalParamTextLang
	EN    SearchOptionalParamTextLang
	VI    SearchOptionalParamTextLang
	ZH    SearchOptionalParamTextLang
	ZH_TW SearchOptionalParamTextLang
}

func GetSearchOptionalParamTextLangEnum() SearchOptionalParamTextLangEnum {
	return SearchOptionalParamTextLangEnum{
		AR: SearchOptionalParamTextLang{
			value: "ar",
		},
		DE: SearchOptionalParamTextLang{
			value: "de",
		},
		RU: SearchOptionalParamTextLang{
			value: "ru",
		},
		FR: SearchOptionalParamTextLang{
			value: "fr",
		},
		KO: SearchOptionalParamTextLang{
			value: "ko",
		},
		PT: SearchOptionalParamTextLang{
			value: "pt",
		},
		JA: SearchOptionalParamTextLang{
			value: "ja",
		},
		TH: SearchOptionalParamTextLang{
			value: "th",
		},
		TR: SearchOptionalParamTextLang{
			value: "tr",
		},
		ES: SearchOptionalParamTextLang{
			value: "es",
		},
		EN: SearchOptionalParamTextLang{
			value: "en",
		},
		VI: SearchOptionalParamTextLang{
			value: "vi",
		},
		ZH: SearchOptionalParamTextLang{
			value: "zh",
		},
		ZH_TW: SearchOptionalParamTextLang{
			value: "zhTW",
		},
	}
}

func (c SearchOptionalParamTextLang) Value() string {
	return c.value
}

func (c SearchOptionalParamTextLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchOptionalParamTextLang) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
