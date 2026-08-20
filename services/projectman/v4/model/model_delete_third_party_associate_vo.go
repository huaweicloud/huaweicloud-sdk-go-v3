package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteThirdPartyAssociateVo 工作项删除外部链接时的参数对象
type DeleteThirdPartyAssociateVo struct {

	// 新关联外部链接时会创建一条数据，该数据的唯一标识ID，可以在查询外部链接接口以及关联外部链接接口响应体中找到。
	Id string `json:"id"`
}

func (o DeleteThirdPartyAssociateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThirdPartyAssociateVo struct{}"
	}

	return strings.Join([]string{"DeleteThirdPartyAssociateVo", string(data)}, " ")
}
