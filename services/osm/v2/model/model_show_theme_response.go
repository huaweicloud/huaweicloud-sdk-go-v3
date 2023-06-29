package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowThemeResponse Response Object
type ShowThemeResponse struct {

	// 关联主题列表
	Themes *[]RelationTheme `json:"themes,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`

	// 主题名称
	ThemeName      *string `json:"theme_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowThemeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowThemeResponse struct{}"
	}

	return strings.Join([]string{"ShowThemeResponse", string(data)}, " ")
}
