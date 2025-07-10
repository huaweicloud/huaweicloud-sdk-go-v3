package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugModelResourceRequest Request Object
type DeleteDrugModelResourceRequest struct {

	// **参数解释**： 盘古药物分子大模型id。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	Id string `json:"id"`
}

func (o DeleteDrugModelResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugModelResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDrugModelResourceRequest", string(data)}, " ")
}
