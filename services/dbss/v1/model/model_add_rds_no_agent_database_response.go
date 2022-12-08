package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddRdsNoAgentDatabaseResponse struct {

	// 添加失败的数据库实例id
	IllegalDbId *[]string `json:"illegal_db_id,omitempty"`

	// 添加成功的数据库实例id
	LegalDbId      *[]string `json:"legal_db_id,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddRdsNoAgentDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRdsNoAgentDatabaseResponse struct{}"
	}

	return strings.Join([]string{"AddRdsNoAgentDatabaseResponse", string(data)}, " ")
}
