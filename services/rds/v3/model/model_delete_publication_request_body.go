package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicationRequestBody （批量）删除发布。
type DeletePublicationRequestBody struct {

	// 删除的发布id列表（需要是同一实例的发布）。
	PublicationIds []string `json:"publication_ids"`
}

func (o DeletePublicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicationRequestBody struct{}"
	}

	return strings.Join([]string{"DeletePublicationRequestBody", string(data)}, " ")
}
