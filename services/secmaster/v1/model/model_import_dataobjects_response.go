package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataobjectsResponse Response Object
type ImportDataobjectsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *CommonDataObjectImportResponseData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ImportDataobjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataobjectsResponse struct{}"
	}

	return strings.Join([]string{"ImportDataobjectsResponse", string(data)}, " ")
}
