package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHibernateTypeResponse Response Object
type ShowHibernateTypeResponse struct {

	// 租户id
	TenantId *string `json:"tenant_id,omitempty"`

	HibernateType *HibernateType `json:"hibernate_type,omitempty"`

	// 休眠关机时长（天）
	ShutdownDays   *int32 `json:"shutdown_days,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHibernateTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHibernateTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowHibernateTypeResponse", string(data)}, " ")
}
