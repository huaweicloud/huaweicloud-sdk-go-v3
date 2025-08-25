package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCcspTenantImagesResponse Response Object
type ListCcspTenantImagesResponse struct {

	// 满足条件的镜像总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 镜像列表
	Result         *[]ImageInfo `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListCcspTenantImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcspTenantImagesResponse struct{}"
	}

	return strings.Join([]string{"ListCcspTenantImagesResponse", string(data)}, " ")
}
