package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVersionAliasesResponse struct {
	// 函数版本别名列表

	Body           *[]ListVersionAliasResult `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListVersionAliasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionAliasesResponse struct{}"
	}

	return strings.Join([]string{"ListVersionAliasesResponse", string(data)}, " ")
}
