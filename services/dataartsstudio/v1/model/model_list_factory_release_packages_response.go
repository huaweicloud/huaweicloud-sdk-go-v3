package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryReleasePackagesResponse Response Object
type ListFactoryReleasePackagesResponse struct {

	// 发布包信息
	Data *[]ListReleasePackagesRespData `json:"data,omitempty"`

	// 发布包个数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFactoryReleasePackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryReleasePackagesResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryReleasePackagesResponse", string(data)}, " ")
}
