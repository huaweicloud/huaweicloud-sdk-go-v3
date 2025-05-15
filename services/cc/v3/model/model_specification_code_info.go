package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecificationCodeInfo struct {

	// 带宽包实例的规格编码。 bandwidth.aftoela：大陆站+国际站南非-拉美东 bandwidth.aftonla：大陆站+国际站南非-拉美北 bandwidth.aftowla：大陆站+国际站南非-拉美西 bandwidth.aptoaf：国际站亚太-南非 bandwidth.aptoap：国际站亚太-亚太 bandwidth.aptoela：大陆站+国际站亚太-拉美东 bandwidth.aptonla：大陆站+国际站亚太-拉美北 bandwidth.aptowla：大陆站+国际站亚太-拉美西 bandwidth.cmtoaf：国际站中国大陆-南非 bandwidth.cmtoap：国际站中国大陆-亚太 bandwidth.cmtocm：国际站中国大陆-中国大陆 bandwidth.cmtoela：大陆站+国际站中国大陆-拉美东 bandwidth.cmtonla：大陆站+国际站中国大陆-拉美北 bandwidth.cmtowla：大陆站+国际站中国大陆-拉美西 bandwidth.elatoela：大陆站+国际站拉美东-拉美东 bandwidth.elatonla：大陆站+国际站拉美东-拉美北 bandwidth.wlatoela：大陆站+国际站拉美西-拉美东 bandwidth.wlatonla：大陆站+国际站拉美西-拉美北 bandwidth.wlatowla：大陆站+国际站拉美西-拉美西
	SpecCode string `json:"spec_code"`

	BillingMode *BillingModeEnum `json:"billing_mode"`

	// 最大带宽。
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// 最小带宽。
	MimBandwidth *int64 `json:"mim_bandwidth,omitempty"`
}

func (o SpecificationCodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecificationCodeInfo struct{}"
	}

	return strings.Join([]string{"SpecificationCodeInfo", string(data)}, " ")
}
