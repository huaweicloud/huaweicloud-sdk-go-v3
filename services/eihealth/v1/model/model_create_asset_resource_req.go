package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAssetResourceReq struct {

	// **参数解释**： 规格编码。 **约束限制**： 不涉及。 **取值范围**： 仅支持ai4ss.basic.model。 **默认取值**： 不涉及
	SpecCode string `json:"spec_code"`
}

func (o CreateAssetResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetResourceReq struct{}"
	}

	return strings.Join([]string{"CreateAssetResourceReq", string(data)}, " ")
}
