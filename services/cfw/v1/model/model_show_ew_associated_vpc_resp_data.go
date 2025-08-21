package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEwAssociatedVpcRespData **参数解释**： 查询VPC间防火墙使用引流VPC返回体 **取值范围**： 不涉及
type ShowEwAssociatedVpcRespData struct {

	// **参数解释**： 使用引流VPC列表 **取值范围**： 不涉及
	InspectionVpcList *[]InspectionVpcInfo `json:"inspection_vpc_list,omitempty"`
}

func (o ShowEwAssociatedVpcRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEwAssociatedVpcRespData struct{}"
	}

	return strings.Join([]string{"ShowEwAssociatedVpcRespData", string(data)}, " ")
}
