package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateThirdPartyAssociateVo 新增关联外部链接参数。
type CreateThirdPartyAssociateVo struct {

	// 工作项下对应的外部链接的名称。
	Title string `json:"title"`

	// 工作项下对应外部链接的地址。
	Url string `json:"url"`
}

func (o CreateThirdPartyAssociateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThirdPartyAssociateVo struct{}"
	}

	return strings.Join([]string{"CreateThirdPartyAssociateVo", string(data)}, " ")
}
