package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceItemResponseData struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **取值范围：** 服务ID。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务名，用户在[创建服务](CreateInferService.xml)时自定义。 **取值范围：** 支持1-128个字符，可以包含字母、汉字、数字、连字符和下划线。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 服务当前状态。 **取值范围：** - DEPLOYING：部署中。 - FAILED：失败。 - STOPPED：停止。 - RUNNING：运行中。 - DELETING：删除中。 - STOPPING：停止中。 - CONCERNING：告警。 - UPGRADING：升级中。 - ERROR：异常。 - INTERRUPTING：中断中。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 服务版本号，以数字及点号组成，形如1.0.1。 **取值范围：** 1.0.0 ~ 99.99.99，长度不超过8位。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 服务版本数量。 **取值范围：** 不涉及。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释：** 服务描述，由用户[创建服务](CreateInferService.xml)时自行填写。 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 推理服务类型。 **约束限制：** 不涉及。 **取值范围：** - REAL_TIME：在线服务。 - ASYNC_REAL_TIME：异步服务。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 部署类型。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 创建服务的用户名。 **取值范围：** 用户名。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 工作空间ID。 **取值范围：** - 0：默认空间ID。 - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml)。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释：** 在线服务认证类型。 **取值范围：** - TOKEN：IAM Token认证。 - API_KEY：API Key认证。 - NONE：无认证。
	AuthType *string `json:"auth_type,omitempty"`

	// **参数解释：** 模型类型。 **取值范围：** - TEXT_GENERATION：文本生成 - IMAGE_UNDERSTANDING：图像理解 - VIDEO_GENERATION：视频生成 - IMAGE_GENERATION：图像生成 - RERANK：重排序 - VECTOR_MODEL：向量模型 - EMBEDDING：Embedding嵌入
	TaskType string `json:"task_type"`

	// **参数解释：** 在线服务标签数据。
	Tags *[]TagsResponse `json:"tags,omitempty"`

	// **参数解释：** 定时停止配置。
	Schedule *[]ScheduleConfigResponse `json:"schedule,omitempty"`

	// **参数解释：** 当服务或者部署被冻结时返回的冻结类型信息。
	FrozenInfos *[]FrozenInfo `json:"frozen_infos,omitempty"`
}

func (o ServiceItemResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItemResponseData struct{}"
	}

	return strings.Join([]string{"ServiceItemResponseData", string(data)}, " ")
}
