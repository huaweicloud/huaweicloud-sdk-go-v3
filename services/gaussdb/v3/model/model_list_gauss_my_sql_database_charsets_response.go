package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlDatabaseCharsetsResponse Response Object
type ListGaussMySqlDatabaseCharsetsResponse struct {

	// 数据库字符集列表
	Charsets       *[]string `json:"charsets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGaussMySqlDatabaseCharsetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseCharsetsResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseCharsetsResponse", string(data)}, " ")
}
