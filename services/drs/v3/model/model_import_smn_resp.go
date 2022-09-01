package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 录入SMN返回体
type ImportSmnResp struct {

	// 任务ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o ImportSmnResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSmnResp struct{}"
	}

	return strings.Join([]string{"ImportSmnResp", string(data)}, " ")
}
