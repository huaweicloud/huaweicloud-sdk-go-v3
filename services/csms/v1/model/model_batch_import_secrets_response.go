package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportSecretsResponse Response Object
type BatchImportSecretsResponse struct {

	// 失败描述
	ErrorList *[]ErrorInfo `json:"error_list,omitempty"`

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 成功条数
	Success *int32 `json:"success,omitempty"`

	// 失败条数
	Failed         *int32 `json:"failed,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchImportSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportSecretsResponse struct{}"
	}

	return strings.Join([]string{"BatchImportSecretsResponse", string(data)}, " ")
}
