package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImageVersionsResponse Response Object
type ListV2TenantImageVersionsResponse struct {

	// **参数解释**：自定义镜像版本总记录数。 **取值范围**：最小值为0，最大值为100000。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：镜像版本列表，每个元素为一个V2ImageVersionInfo对象，包含镜像版本ID、镜像版本、镜像分类、镜像架构、状态等信息。
	ImageVersions  *[]V2ImageVersionInfo `json:"image_versions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListV2TenantImageVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImageVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListV2TenantImageVersionsResponse", string(data)}, " ")
}
