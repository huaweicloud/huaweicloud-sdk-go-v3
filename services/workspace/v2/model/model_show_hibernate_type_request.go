package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHibernateTypeRequest Request Object
type ShowHibernateTypeRequest struct {
}

func (o ShowHibernateTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHibernateTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowHibernateTypeRequest", string(data)}, " ")
}
