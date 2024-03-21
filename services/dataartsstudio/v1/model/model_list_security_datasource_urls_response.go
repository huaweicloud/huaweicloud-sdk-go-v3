package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDatasourceUrlsResponse Response Object
type ListSecurityDatasourceUrlsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// url列表
	Urls           *[]UrlDto `json:"urls,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityDatasourceUrlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceUrlsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceUrlsResponse", string(data)}, " ")
}
