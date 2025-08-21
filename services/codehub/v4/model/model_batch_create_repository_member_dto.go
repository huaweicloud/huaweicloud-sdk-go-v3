package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateRepositoryMemberDto struct {

	// **参数解释：** 批量添加是否成功 **约束限制：** - success，全部添加成功。 - error，未全部添加成功。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 批量添加结果详情。 **约束限制：** 不涉及。
	Result *[]CreateRepositoryMemberResponseDto `json:"result,omitempty"`
}

func (o BatchCreateRepositoryMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRepositoryMemberDto struct{}"
	}

	return strings.Join([]string{"BatchCreateRepositoryMemberDto", string(data)}, " ")
}
