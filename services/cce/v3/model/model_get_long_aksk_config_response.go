package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetLongAkskConfigResponse Response Object
type GetLongAkskConfigResponse struct {

	// **参数解释：** 项目是否上传了LongAKSK。 **约束限制：** 不涉及 **取值范围：** - false: 未上传LongAKSK - true: 已上传LongAKSK  **默认取值：** 不涉及
	HasUploadedLongAKSK *bool `json:"hasUploadedLongAKSK,omitempty"`

	// **参数解释：** 新建集群是否启用LongAKSK。 **约束限制：** 不涉及 **取值范围：** - false: 新建集群不启用LongAKSK - true: 新建集群启用LongAKSK  **默认取值：** 不涉及
	EnableLongAKSKInNewCluster *bool `json:"enableLongAKSKInNewCluster,omitempty"`
	HttpStatusCode             int   `json:"-"`
}

func (o GetLongAkskConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetLongAkskConfigResponse struct{}"
	}

	return strings.Join([]string{"GetLongAkskConfigResponse", string(data)}, " ")
}
