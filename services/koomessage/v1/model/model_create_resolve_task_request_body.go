package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolveTaskRequestBody 生成短链请求体。
type CreateResolveTaskRequestBody struct {

	// 智能信息模板ID，由9位数字组成。
	TplId string `json:"tpl_id"`

	// 短信签名列表，需要与最终发送短信的签名一致，才能解析。  > 最多传入5个签名，单个签名内容不包括中括号。
	SmsSigns []string `json:"sms_signs"`

	// 短链最大解析次数。  >个性化短链只支持最大解析数为1，设置其他值无效。 >群发短链设置解析次数需要大于等于2。
	ResolvingTimes int32 `json:"resolving_times"`

	// 生成短链类型。  - group：群发 - individual：个性化  > 使用动态参数模板时，该字段只能为individual。
	AimCodeType string `json:"aim_code_type"`

	// 生成短码方式。  - 1：标准 - 2：自定义  > 默认1，即标准生成短码。
	GenerationType *string `json:"generation_type,omitempty"`

	// 自定义短链域名，由大小写字母和数字组成的二级域名。  > generation_type为2时，此参数为必填。域名需要提前报备，请联系KooMessage运营人员进行域名报备，域名区分生成短码方式，如报备的是标准生成短码方式，则在自定义生成短码时不能使用此域名。
	Domain *string `json:"domain,omitempty"`

	// 失效时间（天）。aim_code_type为group时，取值范围为1~180；aim_code_type为individual个性化时，取值范围为1~7。  > 失效时间精确到秒，例如参数设置为1，创建时间为2022-07-22 21:10:12，过期时间为2022-07-23 21:10:12。
	ExpirationTime int32 `json:"expiration_time"`

	// 短链参数列表。一次请求最多生成100个短链。  > OPPO模板一次最多申请10个短链。 > 群发短链每次调用最多申请一条。
	Params []CreateShortChainParam `json:"params"`
}

func (o CreateResolveTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolveTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResolveTaskRequestBody", string(data)}, " ")
}
