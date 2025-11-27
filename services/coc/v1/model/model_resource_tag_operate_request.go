package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagOperateRequest struct {

	// **参数解释：** 标签键值对。 **取值范围：** 不涉及。
	Tags []TagResponse `json:"tags"`
}

func (o ResourceTagOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagOperateRequest struct{}"
	}

	return strings.Join([]string{"ResourceTagOperateRequest", string(data)}, " ")
}
