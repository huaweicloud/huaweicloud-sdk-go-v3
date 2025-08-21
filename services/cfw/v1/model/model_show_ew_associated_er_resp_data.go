package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEwAssociatedErRespData **参数解释**： 查询VPC间防火墙使用的企业路由器信息返回体 **取值范围**： 不涉及
type ShowEwAssociatedErRespData struct {

	// **参数解释**： 企业路由器列表 **取值范围**： 不涉及
	ErList *[]ErInfo `json:"er_list,omitempty"`
}

func (o ShowEwAssociatedErRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEwAssociatedErRespData struct{}"
	}

	return strings.Join([]string{"ShowEwAssociatedErRespData", string(data)}, " ")
}
