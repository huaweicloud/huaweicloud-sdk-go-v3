package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceShareAssociationReqBody The request body of AssociateResourceShare and DisassociateResourceShare operations.
type ResourceShareAssociationReqBody struct {

	// 指定与资源共享实例关联的一个或多个资源使用者的列表。
	Principals *[]string `json:"principals,omitempty"`

	// 指定与资源共享实例关联的一个或多个共享资源URN的列表。
	ResourceUrns *[]string `json:"resource_urns,omitempty"`
}

func (o ResourceShareAssociationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceShareAssociationReqBody struct{}"
	}

	return strings.Join([]string{"ResourceShareAssociationReqBody", string(data)}, " ")
}
