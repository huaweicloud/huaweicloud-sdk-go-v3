package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeEnterpriseRealnameAuthsReq struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 认证方案： 1：单位证件扫描
	IdentifyType int32 `json:"identify_type"`

	// 单位证件类型： 0：企业营业执照1：事业单位法人证书2：社会团体法人登记证书3：行政执法主体资格证4：组织机构代码证99：其他 此参数不携带或携带值为null时，不赋值。
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// 单位证件认证时证件附件的文件URL。 附件地址必须按照顺序填写，先填写单位证件的附件，如果存在单位人员信息，再填写单位人员的身份证附件。 单位证件顺序为： 第1张单位证件照附件， 单位人员的证件顺序为： 第1张个人身份证的人像面 第2张个人身份证的国徽面 以营业执照举例，假设存在法人的情况下，第1张上传的是营业执照扫描件abc.023，第2张是法人的身份证人像面照片def004，第3张是法人的国徽面照片gh007，那么上传顺序需要是： abc023 def004 gh007 的顺序填写URL（文件名称区分大小写） 证件附件目前仅仅支持jpg、jpeg、bmp、png、gif、pdf格式，单个文件最大不超过10M。 这个URL是相对URL，不需要包含桶名和download目录，只要包含download目录下的子目录和对应文件名称即可。举例如下： 如果上传的证件附件在桶中的位置是： https://bucketname.obs.Endpoint.myhuaweicloud.com/download/abc023.jpg，该字段填写abc023.jpg； 如果上传的证件附件在桶中的位置是： https://bucketname.obs.Endpoint.myhuaweicloud.com/download/test/abc023.jpg，该字段填写test/abc023.jpg。
	VerifiedFileUrl []string `json:"verified_file_url"`

	// 单位名称。 不能全是数字、特殊字符、空格。
	CorpName string `json:"corp_name"`

	// 单位证件号码。
	VerifiedNumber string `json:"verified_number"`

	// 实名认证填写的注册国家。国家的两位字母简码。 例如：注册国家为“中国”请填写“CN”。 此参数不携带或携带值为null时，不赋值；携带值为空串时，赋值空串。
	RegCountry *string `json:"reg_country,omitempty"`

	// 实名认证企业注册地址。 此参数不携带或携带值为null时，不赋值；携带值为空串时，赋值空串。
	RegAddress *string `json:"reg_address,omitempty"`

	// 变更类型： 1：个人变企业
	ChangeType int32 `json:"change_type"`

	// 华为分给合作伙伴的平台标识。 该标识的具体值由华为分配。获取方法请参见如何获取xaccountType的取值。
	XaccountType string `json:"xaccount_type"`

	EnterprisePerson *EnterprisePersonNew `json:"enterprise_person,omitempty"`
}

func (o ChangeEnterpriseRealnameAuthsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEnterpriseRealnameAuthsReq struct{}"
	}

	return strings.Join([]string{"ChangeEnterpriseRealnameAuthsReq", string(data)}, " ")
}
