package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeImageWhiteListRequestBody struct {

	// **参数解释**： 白名单规则类型 **约束限制**: 不涉及 **取值范围**: - all_images：白名单应用于全部镜像。 - specific_image_types：白名单应用于指定镜像类型(仅用于指定仓库镜像类型)。 - specific_images：白名单应用于指定镜像。  **默认取值**: 不涉及
	RuleType *string `json:"rule_type,omitempty"`

	QueryInfo *AddImageWhiteListsRequestBodyQueryInfo `json:"query_info,omitempty"`

	// 指定具体镜像
	ImageInfo *[]AddImageWhiteListsRequestBodyImageInfo `json:"image_info,omitempty"`

	// 白名单的描述信息
	Description *string `json:"description,omitempty"`
}

func (o ChangeImageWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeImageWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeImageWhiteListRequestBody", string(data)}, " ")
}
