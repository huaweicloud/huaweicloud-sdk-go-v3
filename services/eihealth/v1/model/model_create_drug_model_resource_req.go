package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDrugModelResourceReq struct {

	// **参数解释**： 规格编码。 **约束限制**： 不涉及。 **取值范围**： 仅支持eihealth.platform.pkg.spec.drug.profession。 **默认取值**： 不涉及
	SpecCode string `json:"spec_code"`
}

func (o CreateDrugModelResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugModelResourceReq struct{}"
	}

	return strings.Join([]string{"CreateDrugModelResourceReq", string(data)}, " ")
}
