package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageCachesRequest Request Object
type ListImageCachesRequest struct {

	// **参数解释：** 按单个镜像缓存名称进行过滤，不支持模糊匹配。 **约束限制：** 不涉及 **取值范围：** 以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围1-128位，且不能以中划线(-)结尾。 **默认取值：** 无
	Name *string `json:"name,omitempty"`
}

func (o ListImageCachesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageCachesRequest struct{}"
	}

	return strings.Join([]string{"ListImageCachesRequest", string(data)}, " ")
}
