package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateThirdPartyAssociateVo 工作项更新外部链接时的参数对象
type UpdateThirdPartyAssociateVo struct {

	// 工作项下关联外部链接的名称。
	Title string `json:"title"`

	// 工作项下关联外部链接的地址。
	Url string `json:"url"`

	// 新关联外部链接时会创建一条数据，该数据的唯一标识ID，可以在查询外部链接接口以及关联外部链接接口响应体中找到。
	Id string `json:"id"`
}

func (o UpdateThirdPartyAssociateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateThirdPartyAssociateVo struct{}"
	}

	return strings.Join([]string{"UpdateThirdPartyAssociateVo", string(data)}, " ")
}
