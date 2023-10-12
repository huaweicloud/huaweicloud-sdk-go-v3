package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppGroupAuthorizeReq 为应用(组)添加(取消)授权的操作，必须为app_group_ids,app_ids同时赋值。 > - 批量操作的应用组或者发布应用的授权模式必须统一，即要么都为应用类型的授权，要么都为应用组类型的授权，否则会直接失败(授权类型在创建应用组时指定)。
type AppGroupAuthorizeReq struct {

	// 应用组ID,最多同时操作10个
	AppGroupIds []string `json:"app_group_ids"`

	// 用户(组),单次最多允许操作50个用户(组)
	Accounts []AccountInfo `json:"accounts"`
}

func (o AppGroupAuthorizeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppGroupAuthorizeReq struct{}"
	}

	return strings.Join([]string{"AppGroupAuthorizeReq", string(data)}, " ")
}
