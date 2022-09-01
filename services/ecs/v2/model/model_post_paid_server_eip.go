package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostPaidServerEip struct {

	// 弹性IP地址类型。  详情请参见“[申请弹性公网IP](https://support.huaweicloud.com/api-eip/eip_api_0001.html)”章节的“publicip”字段说明。
	Iptype string `json:"iptype" xml:"iptype"`

	Bandwidth *PostPaidServerEipBandwidth `json:"bandwidth" xml:"bandwidth"`

	Extendparam *PostPaidServerEipExtendParam `json:"extendparam,omitempty" xml:"extendparam"`
}

func (o PostPaidServerEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEip struct{}"
	}

	return strings.Join([]string{"PostPaidServerEip", string(data)}, " ")
}
