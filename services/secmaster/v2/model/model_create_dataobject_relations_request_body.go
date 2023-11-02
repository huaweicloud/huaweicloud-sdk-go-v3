package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectRelationsRequestBody 关联dataobject列表请求body体
type CreateDataobjectRelationsRequestBody struct {

	// 关联dataobject的ID列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o CreateDataobjectRelationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectRelationsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataobjectRelationsRequestBody", string(data)}, " ")
}
