package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiInfoRequest Request Object
type ShowApiInfoRequest struct {

	// **参数解释：** 待查询版本号。 **约束限制：** 不涉及。 **取值范围：** v2。 **默认取值：** 不涉及。
	Version string `json:"version"`
}

func (o ShowApiInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowApiInfoRequest", string(data)}, " ")
}
