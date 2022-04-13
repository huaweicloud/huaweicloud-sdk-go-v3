package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSharedTagRequest struct {
	// MIME类型

	ContentType string `json:"Content-Type"`
	// 共享ID

	ShareId string `json:"share_id"`
	// 标签的键,最大长度36个字符。  key不能为空，不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。只能包含大写字母、小写字母、数字，特殊字符\"-\"和\"_\"。  说明：调用删除共享标签接口删除标签时，如果标签的键中存在不被URL直接解析的特殊字符，需要对标签的键进行URL转义处理。

	Key string `json:"key"`
}

func (o DeleteSharedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSharedTagRequest", string(data)}, " ")
}
