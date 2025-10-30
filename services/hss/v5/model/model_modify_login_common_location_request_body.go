package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginCommonLocationRequestBody 国家列表
type ModifyLoginCommonLocationRequestBody struct {

	// **参数解释**： 国家城市的编码 **取值范围**： 字符长度1-32位
	AreaCode *int32 `json:"area_code,omitempty"`

	// **参数解释**: 服务器列表 **取值范围**: 最小值0，最大值100
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ModifyLoginCommonLocationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginCommonLocationRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyLoginCommonLocationRequestBody", string(data)}, " ")
}
