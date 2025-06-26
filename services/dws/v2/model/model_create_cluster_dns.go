package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterDns **参数解释**： 申请的域名信息。 **取值范围**： 不涉及。
type CreateClusterDns struct {

	// **参数解释**： 待创建的域名。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 域名类型。 **取值范围**： public：公网域名。 private：内网域名。
	Type string `json:"type"`

	// **参数解释**： 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。默认值为300s。 **取值范围**： 300~2147483647。
	Ttl int32 `json:"ttl"`
}

func (o CreateClusterDns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDns struct{}"
	}

	return strings.Join([]string{"CreateClusterDns", string(data)}, " ")
}
