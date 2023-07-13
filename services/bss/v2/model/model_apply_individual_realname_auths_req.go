package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyIndividualRealnameAuthsReq struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 认证方案： 0：个人证件认证 4：个人银行卡认证
	IdentifyType int32 `json:"identify_type"`

	// 证件类型： 0：身份证，上传的附件为3张，第1张是身份证人像面，第2张是身份证国徽面，第3张是个人手持身份证人像面； 3：护照，上传的附件为3张，第1张是护照个人资料页，第2张是，护照入境盖章页，第3张是手持护照个人资料页； 5：港澳通行证，上传的附件为3张，第1张是港澳居民来往内地通行证正面（人像面），第2张是港澳居民来往内地通行证反面，第3张是手持港澳居民来往内地通行证人像面； 6：台湾通行证，上传的附件为3张，第1张是台湾居民来往大陆通行证正面（人像面），第2张是台湾居民来往大陆通行证反面，第3张是手持台湾居民来往大陆通行证人像面； 9：港澳居民居住证，上传的附件为3张，第1张是港澳居民居住证人像面，第2张是，港澳居民居住证国徽面，第3张是手持港澳居民居住证人像面照片； 10：台湾居民居住证，上传的附件为3张，第1张是台湾居民居住证人像面，第2张是台湾居民居住证国徽面，第3张是手持台湾居民居住证人像面照片。 当identify_type=0的时候，该字段需要填写，否则忽略该字段的取值。
	VerifiedType *int32 `json:"verified_type,omitempty"`

	// 个人证件认证时证件附件的文件URL，该URL地址必须按照顺序填写。以身份证举例，譬如身份证人像面文件名称是abc023，国徽面是def004，个人手持身份证人像面是gh007，那么这个地方需要按照 abc023 def004 gh007 的顺序填写URL（文件名称区分大小写）。 个人银行卡认证时直接上传一张个人扫脸的图片附件即可。证件附件目前仅仅支持jpg、jpeg、bmp、png、gif、pdf格式，单个文件最大不超过10M。这个URL是相对URL，不需要包含桶名和download目录，只要包含download目录下的子目录和对应文件名称即可。举例如下：如果上传的证件附件在桶中的位置是：https://bucketname.obs.Endpoint.myhuaweicloud.com/download/abc023.jpg，该字段填写abc023.jpg； 如果上传的证件附件在桶中的位置是：https://bucketname.obs.Endpoint.myhuaweicloud.com/download/test/abc023.jpg，该字段填写test/abc023.jpg。
	VerifiedFileUrl []string `json:"verified_file_url"`

	// 姓名。
	Name string `json:"name"`

	// 证件号码。
	VerifiedNumber string `json:"verified_number"`

	// 变更类型： -1：首次实名认证
	ChangeType *int32 `json:"change_type,omitempty"`

	// 华为分给合作伙伴的平台标识。 该标识的具体值由华为分配。获取方法请参见如何获取xaccountType的取值。
	XaccountType string `json:"xaccount_type"`

	BankCardInfo *BankCardInfoV2 `json:"bank_card_info,omitempty"`
}

func (o ApplyIndividualRealnameAuthsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyIndividualRealnameAuthsReq struct{}"
	}

	return strings.Join([]string{"ApplyIndividualRealnameAuthsReq", string(data)}, " ")
}
