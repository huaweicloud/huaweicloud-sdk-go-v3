package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectRelationRequestBody 关联数据对象列表请求body体
type CreateDataobjectRelationRequestBody struct {

	// 关联数据对象的ID列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o CreateDataobjectRelationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectRelationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataobjectRelationRequestBody", string(data)}, " ")
}
