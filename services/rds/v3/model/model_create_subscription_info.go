package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionInfo 订阅详情
type CreateSubscriptionInfo struct {

	// 订阅数据库名。
	SubscriptionDatabase string `json:"subscription_database"`

	// 订阅方式，push:推送。
	SubscriptionType string `json:"subscription_type"`

	// 初始化类型：不初始化(do_not)、立即初始化(immediate)、首次同步初始化(at_first_sync)。
	InitializeAt string `json:"initialize_at"`

	InitializeInfo *CreateSubscriptionInfoInitializeInfo `json:"initialize_info,omitempty"`

	// 独立的分发代理。  true：使用。 false：不使用。
	IndependentAgent *bool `json:"independent_agent,omitempty"`

	JobSchedule *OperateUsedJobSchedule `json:"job_schedule"`

	// 备份文件名称。若该值不为空，则订阅初始化方式为通过备份文件初始化。
	BakFileName *string `json:"bak_file_name,omitempty"`

	// 备份文件所在的obs桶名。 若bak_file_name为空，该参数无效。 若该值为空，则备份文件来源为rds的备份文件。 若该值不为空，则备份文件来源为用户obs桶。
	BakBucketName *string `json:"bak_bucket_name,omitempty"`

	PublicationSubscription *CreateSubscription4PublicationInfo `json:"publication_subscription,omitempty"`

	LocalSubscription *CreateSubscription4InstanceInfo `json:"local_subscription,omitempty"`
}

func (o CreateSubscriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionInfo", string(data)}, " ")
}
