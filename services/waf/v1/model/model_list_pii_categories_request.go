package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPiiCategoriesRequest Request Object
type ListPiiCategoriesRequest struct {

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** en-us
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListPiiCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPiiCategoriesRequest struct{}"
	}

	return strings.Join([]string{"ListPiiCategoriesRequest", string(data)}, " ")
}
