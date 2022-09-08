package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RateOnDemandReq struct {

	// 项目ID。  说明： 使用客户Token，可以调用[通过assume_role方式获取用户token](https://support.huaweicloud.com/api-iam/iam_30_0003.html)接口获取“regionId”的取值对应的“project id”。具体请参见[如何将合作伙伴Token置换为客户Token](https://support.huaweicloud.com/bpconsole_faq/bpconsole_apifaq_00001.html)的步骤2。IAM子用户调用此接口，需要IAM主账号授权，具体请参考[创建用户组并授权](https://support.huaweicloud.com/usermanual-iam/iam_03_0001.html)。
	ProjectId string `json:"project_id"`

	// 查询价格结果的精度模式。 0：询价结果默认精度截取，即最长保留到元后6位小数点，如0.000001元1：询价结果保留10位精度，即最长保留到元后10位小数点，如1.0000000001元  说明： 如果询价结果只到元后2位或者3位，那么价格也只到元后2位或者3位，不管传0或者传1都一样，只有询价结果到了小数点后面6位以上，传0和传1才有区别。
	InquiryPrecision *int32 `json:"inquiry_precision,omitempty"`

	// 产品信息列表，询价时要询价产品的信息的列表，具体参见表1。
	ProductInfos []DemandProductInfo `json:"product_infos"`
}

func (o RateOnDemandReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RateOnDemandReq struct{}"
	}

	return strings.Join([]string{"RateOnDemandReq", string(data)}, " ")
}
