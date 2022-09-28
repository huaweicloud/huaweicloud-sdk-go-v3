package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type HealthCodeResult struct {

	// 防疫码类别： - 健康码：health_code - 核酸检测记录：pcr_test_record - 通信行程卡：travel_card - 其他：other
	Type string `json:"type"`

	// 姓名
	Name string `json:"name"`

	// 身份证号码
	IdcardNumber string `json:"idcard_number"`

	// 手机号码
	PhoneNumber string `json:"phone_number"`

	// 省份
	Province string `json:"province"`

	// 城市
	City string `json:"city"`

	// 健康码或行程卡的更新时间
	Time string `json:"time"`

	// 健康码或行程卡颜色。 健康码颜色可选值包括：  - \"green\"，绿码 - \"yellow\"，黄码 - \"red\"，红码 - \"gray\"，灰码  行程卡颜色可选值包括：  - \"green\"，绿码 - \"yellow\"，黄码 - \"red\"，红码
	Color string `json:"color"`

	// 疫苗接种情况
	VaccinationStatus string `json:"vaccination_status"`

	// 核酸检测结果，可选值包括： - \"positive\",即阳性 - \"negative\",即阴性 - \"unknown\",未知
	PcrTestResult string `json:"pcr_test_result"`

	// 核酸检测机构
	PcrTestOrganization string `json:"pcr_test_organization"`

	// 核酸检测结果更新时间
	PcrTestTime string `json:"pcr_test_time"`

	// 核酸检测采样时间
	PcrSamplingTime string `json:"pcr_sampling_time"`

	// 行程卡的途径地址
	ReachedCity []string `json:"reached_city"`

	// 各个字段的置信度。
	Confidence *interface{} `json:"confidence"`

	// 代表检测识别出来的文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList []HealthCodeWordsBlockList `json:"words_block_list"`
}

func (o HealthCodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCodeResult struct{}"
	}

	return strings.Join([]string{"HealthCodeResult", string(data)}, " ")
}
