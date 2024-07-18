package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimResolveTask 生成发送任务短链配置信息
type AimResolveTask struct {

	// 智能信息模板ID，由9位数字组成。
	TplId string `json:"tpl_id"`

	// 短链的最大解析次数。  > 个性化短链只支持最大解析数为1，设置其他值无效
	ResolveTimes int32 `json:"resolve_times"`

	// 智能信息编码类型。 - individual：个性化
	AimCodeType string `json:"aim_code_type"`

	// 生成短码方式。  - 1：标准 - 2：自定义  > 默认1，即标准生成短码。
	GenerationType *string `json:"generation_type,omitempty"`

	// 自定义短链域名，由大小写字母和数字组成的二级域名。   > 自定义短码即generation_type为2时，此参数为必填。域名需要提前报备，请联系KooMessage运营人员进行域名报备，域名区分生成短码方式，如报备的是标准生成短码方式，则在自定义生成短码时不能使用此域名。
	Domain *string `json:"domain,omitempty"`

	// 失效时间（天）。aim_code_type为individual个性化时，取值范围为1~7。  > 失效时间精确到秒，例如参数设置为1，创建时间为2022-07-22 21:10:12，过期时间为2022-07-23 21:10:12。
	ExpirationTime int32 `json:"expiration_time"`

	// 短链解析详情列表。一次请求最多100个短链。
	Params []CreateResolveTaskParam `json:"params"`
}

func (o AimResolveTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimResolveTask struct{}"
	}

	return strings.Join([]string{"AimResolveTask", string(data)}, " ")
}
