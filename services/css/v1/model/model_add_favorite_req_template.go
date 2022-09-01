package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddFavoriteReqTemplate struct {

	// 模板名称。
	TemplateName string `json:"templateName" xml:"templateName"`

	// 模板描述。
	Desc *string `json:"desc,omitempty" xml:"desc"`
}

func (o AddFavoriteReqTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavoriteReqTemplate struct{}"
	}

	return strings.Join([]string{"AddFavoriteReqTemplate", string(data)}, " ")
}
