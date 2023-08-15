package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportSmnResp 录入SMN返回体
type ImportSmnResp struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ImportSmnResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSmnResp struct{}"
	}

	return strings.Join([]string{"ImportSmnResp", string(data)}, " ")
}
