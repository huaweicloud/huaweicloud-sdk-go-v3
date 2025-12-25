package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetrieveScriptsResponse Response Object
type ListRetrieveScriptsResponse struct {

	// 计数
	Listcount *int64 `json:"listcount,omitempty"`

	// 检索
	Records        *[]RetrieveScript `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListRetrieveScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetrieveScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListRetrieveScriptsResponse", string(data)}, " ")
}
