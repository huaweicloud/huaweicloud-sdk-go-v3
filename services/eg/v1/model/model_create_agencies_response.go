package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAgenciesResponse struct {

	// 服务委托授权记录。
	Agencies       *[]string `json:"agencies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateAgenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesResponse struct{}"
	}

	return strings.Join([]string{"CreateAgenciesResponse", string(data)}, " ")
}
