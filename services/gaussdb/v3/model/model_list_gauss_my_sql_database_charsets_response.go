package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGaussMySqlDatabaseCharsetsResponse struct {

	// 数据库字符集列表
	Charsets       *[]string `json:"charsets,omitempty" xml:"charsets"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGaussMySqlDatabaseCharsetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseCharsetsResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseCharsetsResponse", string(data)}, " ")
}
