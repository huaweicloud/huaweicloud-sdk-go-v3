package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterNameReq **参数解释**： 修改集群名称请求。 **约束限制**： guestAgent插件版本8.3.1及以上才支持。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterNameReq struct {

	// **参数解释**： 修改集群名称请求。 **约束限制**： guestAgent插件版本8.3.1及以上才支持。 **取值范围**： 中文或英文字母开头，3~64位长度。 **默认取值**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o ClusterNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNameReq struct{}"
	}

	return strings.Join([]string{"ClusterNameReq", string(data)}, " ")
}
