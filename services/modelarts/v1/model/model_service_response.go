package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceResponse **参数解释：** 在线服务返回体
type ServiceResponse struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务名，用户在[创建服务](CreateInferService.xml)时自定义的名称。 **取值范围：** 支持1-64位字符，可包含字母、中文、数字、中划线、下划线。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 工作空间ID。 **取值范围：** - 0：默认空间ID。 - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml)。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 推理服务类型。 **取值范围：** - REAL_TIME：在线服务。 - ASYNC_REAL_TIME：异步服务。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 部署方式。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 服务当前状态。 **取值范围：** - DEPLOYING：部署中 。 - FAILED：失败 。 - STOPPED：停止。 - RUNNING：运行中。 - DELETING：删除中 。 - STOPPING：停止中 。 - CONCERNING：告警 。 - UPGRADING：升级中 。 - ERROR：异常 。 - INIT：待部署。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 在线服务失败原因。 **取值范围：** 不涉及。
	FailureReason string `json:"failure_reason"`

	Version *ServiceVersionResponse `json:"version,omitempty"`

	// **参数解释：** 在线服务访问地址，创建服务接口无返回，详情接口中返回。
	PredictUrl *[]PredictUrlResponse `json:"predict_url,omitempty"`

	// **参数解释：** 服务绑定的dispatcher组ID，创建服务接口无返回，详情接口中返回。 **取值范围：** 不涉及。
	DispatcherGroupId *string `json:"dispatcher_group_id,omitempty"`

	// **参数解释：** TMS对接标签类。
	Tags *[]TagsResponse `json:"tags,omitempty"`

	// **参数解释：** 部署超时时间。
	DeployTimeoutMinutes *int32 `json:"deploy_timeout_minutes,omitempty"`

	// **参数解释：** 定时停止配置。
	Schedule *[]ScheduleConfigResponse `json:"schedule,omitempty"`

	// **参数解释：** 创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字，如1609459200000。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字，如1609459200000。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释：** 模型类型。 **取值范围：** - TEXT_GENERATION：文本生成 - IMAGE_UNDERSTANDING：图像理解 - VIDEO_GENERATION：视频生成 - IMAGE_GENERATION：图像生成 - RERANK：重排序 - VECTOR_MODEL：向量模型 - EMBEDDING：Embedding嵌入
	TaskType *string `json:"task_type,omitempty"`

	// **参数解释：** 服务提供者的账号id（创建服务时通过X-Auth-Token-Provider请求头解析iam token而来），该值不为空时，该服务的所有更新操作需要在请求头中添加X-Auth-Token-Provider，取值为该账号id的domain级或project级token。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 当服务或者部署被冻结时返回的冻结类型信息。
	FrozenInfos *[]FrozenInfo `json:"frozen_infos,omitempty"`
}

func (o ServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceResponse struct{}"
	}

	return strings.Join([]string{"ServiceResponse", string(data)}, " ")
}
