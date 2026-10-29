package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskBasicInfoV3 **参数解释：** 任务信息。
type TaskBasicInfoV3 struct {

	// **参数解释：** 头像颜色。 **取值范围：** - pure-blue：蓝色。 - pure-purple：紫色。 - pure-tomato：红色。 - pure-green：绿色。 - pure-orange：橙色。 - pure-pink：粉色。
	AvastarType *string `json:"avastarType,omitempty"`

	// **参数解释：** 分支名称。 **取值范围：** 长度为[1~200]个字符。
	BranchName *string `json:"branchName,omitempty"`

	// **参数解释：** 是否可以更新忽略文件模式。 **取值范围：** - true：能更新忽略文件模式。 - false：不能更新忽略文件模式。
	CanChangeIgnoreMode *bool `json:"canChangeIgnoreMode,omitempty"`

	// **参数解释：** 是否有删除任务的权限。 **取值范围：** - true：有删除的权限。 - false：没有删除的权限。
	CanDelete *bool `json:"canDelete,omitempty"`

	// **参数解释：** 是否有编辑的权限。 **取值范围：** - true：有编辑的权限。 - false：没有编辑的权限。
	CanEdit *bool `json:"canEdit,omitempty"`

	// **参数解释：** 是否有切换任务模板的权限。 **取值范围：** - true：有切换任务模板的权限。 - false：没有切换任务模板的权限。
	CanEditTemplateSwitch *bool `json:"canEditTemplateSwitch,omitempty"`

	// **参数解释：** 是否有执行的权限。 **取值范围：** - true：有执行的权限。 - false：没有执行的权限。
	CanExecute *bool `json:"canExecute,omitempty"`

	// **参数解释：** 是否有读的权限。 **取值范围：** - true：有读的权限。 - false：没有读的权限。
	CanRead *bool `json:"canRead,omitempty"`

	// **参数解释：** 自定义参数列表。
	CheckParamList *[]CustomParameter `json:"checkParamList,omitempty"`

	CheckParamMap *TaskBasicInfoV3CheckParamMap `json:"checkParamMap,omitempty"`

	// **参数解释：** 任务的配置模板的ID。 **取值范围：** 长度为[1~128]个字符。
	ConfigTemplateId *string `json:"configTemplateId,omitempty"`

	// **参数解释：** 任务的创建时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：** 规则集列表，该字段一般返回空数组[]。
	CriterionSets *[]ConfigTemplateRuleSetInfo `json:"criterionSets,omitempty"`

	// **参数解释：** 自定义参数列表，该字段一般返回空数组[]。
	CustomParams *[]CustomParameter `json:"customParams,omitempty"`

	// **参数解释：** 默认分支的分支名。 **取值范围：** 长度为[1~200]个字符。
	DefaultBranchName *string `json:"defaultBranchName,omitempty"`

	// **参数解释：** 文件编码。 **取值范围：** - UTF-8：UTF-8格式。 - GBK：GBK格式。
	FileEncoding *string `json:"fileEncoding,omitempty"`

	// **参数解释：** 华为云地域。 **取值范围：** 长度为[1~32]个字符。
	HwCloudRegion *string `json:"hwCloudRegion,omitempty"`

	// **参数解释：** 任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **取值范围：** 长度为[1~128]个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 检查目录。 不配置检查目录时，检查代码仓中所有目录下的代码文件； 配置检查目录时，只检查配置目录中的代码文件。 **取值范围：** 长度为[1~200]个字符。
	IncludePaths *string `json:"includePaths,omitempty"`

	// **参数解释：** 是否打开合并请求时触发。 **取值范围：** - true：打开。 - false：关闭。
	IsSetMrWebhook *bool `json:"isSetMrWebhook,omitempty"`

	// **参数解释：** 是否打开代码提交时触发。 **取值范围：** - true：打开。 - false：关闭。
	IsSetPushWebhook *bool `json:"isSetPushWebhook,omitempty"`

	// **参数解释：** 是否打开合并请求时触发。 **取值范围：** - true：打开。 - false：关闭。
	IsSetWebhook *bool `json:"isSetWebhook,omitempty"`

	// **参数解释：** 是否是任务模板。 **取值范围：** - true：是任务模板。 - false：不是任务模板。
	IsTemplateTask *bool `json:"isTemplateTask,omitempty"`

	// **参数解释：** 执行ID，来自代码检查执行任务接口的返回值，每次执行代码检查任务都有唯一的执行ID。 **取值范围：** 长度为[1~128]个字符。
	JobId *string `json:"jobId,omitempty"`

	// **参数解释:** 超时时间，以分钟为单位，格式固定为\"xxx min\"，其中xxx为整数，0 ＜ xxx ≤ 1200。 **取值范围：** 不涉及。
	JobTimeout *string `json:"jobTimeout,omitempty"`

	// **参数解释：** 最后一次检查完成的任务的执行ID。 **取值范围：** 长度为[1~128]个字符。
	LastFinishedFullJobId *string `json:"lastFinishedFullJobId,omitempty"`

	LastJob *JobBasicInfoV3 `json:"lastJob,omitempty"`

	// **参数解释：** 任务的最后修改人。 **取值范围：** 不涉及。
	LastReviser *string `json:"lastReviser,omitempty"`

	// **参数解释：** 任务的上一次更新时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：** 模型集列表，该字段一般返回空数组[]。
	ModelSets *[]string `json:"modelSets,omitempty"`

	// **参数解释：** mr检查分支门禁策略。 **取值范围：** - 0：白名单。
	MrCheckBranchStrategy *int32 `json:"mrCheckBranchStrategy,omitempty"`

	// **参数解释：** 代码检查任务名。 **取值范围：** 长度为[1~128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 消息通知是否使用项目通知设置的配置项，格式为x,x,x,x,x,x，其中x的取值为0或者1。 x为0时代表使用任务的通知设置，x为1时代表使用项目通知设置。 从左到右依次为站内消息通知、邮件通知、企业微信、钉钉、飞书、通用webhook的消息通知配置。 **取值范围：** 格式为x,x,x,x,x,x，其中x的取值为0或者1。
	NotificationConfig *string `json:"notificationConfig,omitempty"`

	// **参数解释：** 代码检查任务所有者。 **取值范围：** 长度为[1~255]个字符。
	Owner *string `json:"owner,omitempty"`

	// **参数解释：** 代码检查任务所有者的用户ID。 **取值范围：** 长度为[1~64]个字符。
	OwnerId *string `json:"ownerId,omitempty"`

	// **参数解释：** 项目的32位UUID，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **取值范围：** 长度为[1~32]个字符。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释：** 项目名。 **取值范围：** 长度为[1~128]个字符。
	ProjectName *string `json:"projectName,omitempty"`

	// **参数解释：** 地域，此字段固定为\"suzhou\"。 **取值范围：** suzhou。
	Region *string `json:"region,omitempty"`

	// **参数解释：** 代码仓地址对应的CodeartsRepo仓的RepoId。 **取值范围：** 长度为[1~128]个字符。
	RepoId *string `json:"repoId,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 长度为[1~1024]个字符。
	RepoUrl *string `json:"repoUrl,omitempty"`

	// **参数解释：** 门禁配置，该字段一般返回空数组[]。
	ReviewData *[]string `json:"reviewData,omitempty"`

	// **参数解释：** 规则集配置，该字段一般返回空数组[]。
	RuleSets *[]string `json:"ruleSets,omitempty"`

	// **参数解释：** 任务来源，一般为\"customer\"。 **取值范围：** customer。
	Source *string `json:"source,omitempty"`

	// **参数解释：** 任务分组的ID。 **取值范围：** 长度为[1~128]个字符。
	TaskGroupId *string `json:"taskGroupId,omitempty"`

	// **参数解释：** 任务ID列表，该字段一般返回空数组[]。
	TaskIds *[]string `json:"taskIds,omitempty"`

	// **参数解释：** 任务所属的租户的租户ID，租户ID的值来源于IAM的租户ID。 **取值范围：** 长度为[1~32]个字符。
	TenantId *string `json:"tenantId,omitempty"`

	// **参数解释：** 查询任务信息是否使用缓存。 **取值范围：** - true：使用缓存。 - false：不使用缓存。
	UseCache *bool `json:"useCache,omitempty"`

	// **参数解释：** 消息通知是否使用项目通知设置的配置项，来自notificationConfig的对应的二进制的取反转化成十进制。 例如，notificationConfig配置为1,1,0,1,0,0，则二进制为110100，取反为001011，转化为10进制时该值为11，最终useWorkspaceNotificationConfig的值为11。 **取值范围：** 0-31
	UseWorkspaceNotificationConfig *int32 `json:"useWorkspaceNotificationConfig,omitempty"`

	// **参数解释：** 版本升级缓冲期，该字段一般返回空数组[]。
	VersionUpdateDelayDays *[]string `json:"versionUpdateDelayDays,omitempty"`

	// **参数解释：** webhook检查模式，一般为0，代表禁止来自其它代码仓的MR。 **取值范围：** 0
	WebhookCheckMode *int32 `json:"webhookCheckMode,omitempty"`

	// **参数解释：** 区域，一般为green。 **取值范围：** green。
	Zone *string `json:"zone,omitempty"`
}

func (o TaskBasicInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicInfoV3 struct{}"
	}

	return strings.Join([]string{"TaskBasicInfoV3", string(data)}, " ")
}
