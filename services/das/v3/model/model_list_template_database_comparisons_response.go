package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateDatabaseComparisonsResponse Response Object
type ListTemplateDatabaseComparisonsResponse struct {

	// 数据库列表
	DbNameList     *[]string `json:"db_name_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTemplateDatabaseComparisonsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateDatabaseComparisonsResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateDatabaseComparisonsResponse", string(data)}, " ")
}
