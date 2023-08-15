package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesResponse Response Object
type ListAgenciesResponse struct {

	// 服务委托授权记录。
	Agencies       *[]string `json:"agencies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAgenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesResponse struct{}"
	}

	return strings.Join([]string{"ListAgenciesResponse", string(data)}, " ")
}
