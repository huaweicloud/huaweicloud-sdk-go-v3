package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRepositoryMemberBodyDto 批量添加成员
type BatchCreateRepositoryMemberBodyDto struct {

	// **参数解释：** 成员信息列表。
	Users []CreateRepositoryMemberDto `json:"users"`
}

func (o BatchCreateRepositoryMemberBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRepositoryMemberBodyDto struct{}"
	}

	return strings.Join([]string{"BatchCreateRepositoryMemberBodyDto", string(data)}, " ")
}
