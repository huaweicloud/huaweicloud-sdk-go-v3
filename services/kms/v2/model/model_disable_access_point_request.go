package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAccessPointRequest Request Object
type DisableAccessPointRequest struct {

	// **参数解释：** 接入点ID。 **约束限制：** 不涉及 **取值范围：** UUID格式，字符长度36-36。 **默认取值：** 不涉及
	AccessPointId string `json:"access_point_id"`
}

func (o DisableAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAccessPointRequest struct{}"
	}

	return strings.Join([]string{"DisableAccessPointRequest", string(data)}, " ")
}
