package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImagesResponse Response Object
type ListV2TenantImagesResponse struct {

	// **参数解释**：符合条件的总数。 **取值范围**：最小值为0，最大值为10000。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：自定义镜像列表，列表中每个元素为V2ImageInfo对象，数组长度不超过10000。 **约束限制**：不涉及。
	Images *[]V2ImageInfo `json:"images,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListV2TenantImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImagesResponse struct{}"
	}

	return strings.Join([]string{"ListV2TenantImagesResponse", string(data)}, " ")
}
