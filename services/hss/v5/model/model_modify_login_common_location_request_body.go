package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginCommonLocationRequestBody 国家列表
type ModifyLoginCommonLocationRequestBody struct {

	// 国家城市的编码
	AreaCode *int32 `json:"area_code,omitempty"`

	// 服务器列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ModifyLoginCommonLocationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginCommonLocationRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyLoginCommonLocationRequestBody", string(data)}, " ")
}
