package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterLongAkskConfigResponse Response Object
type GetClusterLongAkskConfigResponse struct {

	// **参数解释：** 集群是否上传了LongAKSK。 **约束限制：** 不涉及 **取值范围：** - false: 未上传LongAKSK - true: 已上传LongAKSK  **默认取值：** 不涉及
	HasUploadedLongAKSK *bool `json:"hasUploadedLongAKSK,omitempty"`

	// **参数解释：** 是否启用LongAKSK，启用后在集群kube-system命名空间下会创建名称为paas.longaksk的密钥，关闭则会删除该密钥。 **约束限制：** 不涉及 **取值范围：** - false: 禁用LongAKSK - true: 启用LongAKSK  **默认取值：** 不涉及
	EnableLongAKSK *bool `json:"enableLongAKSK,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o GetClusterLongAkskConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterLongAkskConfigResponse struct{}"
	}

	return strings.Join([]string{"GetClusterLongAkskConfigResponse", string(data)}, " ")
}
