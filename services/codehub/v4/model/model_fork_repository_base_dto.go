package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkRepositoryBaseDto struct {

	// **参数解释：** 仓库ID。 **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 是否归档。 **约束限制：** 不涉及。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 产品ID。 **约束限制：** 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释：** 产品名称。 **约束限制：** 不涉及。
	ProductName *string `json:"product_name,omitempty"`
}

func (o ForkRepositoryBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkRepositoryBaseDto struct{}"
	}

	return strings.Join([]string{"ForkRepositoryBaseDto", string(data)}, " ")
}
