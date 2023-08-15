package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunConstituencyParserResponse Response Object
type RunConstituencyParserResponse struct {

	// 成分句法分析结果，用嵌套括号的形式表示一棵树。括号内的第一个元素为子树的标签，若是叶子节点则用符号_代替。第二个元素为子树，若是叶子节点则为字符串。
	Tree *string `json:"tree,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunConstituencyParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunConstituencyParserResponse struct{}"
	}

	return strings.Join([]string{"RunConstituencyParserResponse", string(data)}, " ")
}
