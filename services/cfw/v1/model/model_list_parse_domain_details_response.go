package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParseDomainDetailsResponse Response Object
type ListParseDomainDetailsResponse struct {

	// 域名id列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListParseDomainDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParseDomainDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListParseDomainDetailsResponse", string(data)}, " ")
}
