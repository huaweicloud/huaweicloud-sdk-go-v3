package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesV5Response Response Object
type ListAgenciesV5Response struct {

	// 委托及信任委托列表。
	Agencies *[]Agency `json:"agencies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAgenciesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesV5Response struct{}"
	}

	return strings.Join([]string{"ListAgenciesV5Response", string(data)}, " ")
}
