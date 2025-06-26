package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterDns **参数解释**： 修改的域名信息。 **取值范围**： 不涉及。
type ModifyClusterDns struct {

	// **参数解释**： 待修改的域名。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 域名类型。 **取值范围**： public：公网域名。 private：内网域名。
	Type string `json:"type"`

	// **参数解释**： 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。默认值为300s。 **取值范围**： 300~2147483647。
	Ttl int32 `json:"ttl"`
}

func (o ModifyClusterDns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterDns struct{}"
	}

	return strings.Join([]string{"ModifyClusterDns", string(data)}, " ")
}
