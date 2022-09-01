package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGaussMySqlInstanceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *MysqlInstanceRequest `json:"body,omitempty" xml:"body"`
}

func (o CreateGaussMySqlInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlInstanceRequest", string(data)}, " ")
}
