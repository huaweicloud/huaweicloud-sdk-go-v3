package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeLabels **参数解释**：节点的标签信息。
type NodeLabels struct {

	// **参数解释**：节点所在的集群名称。 **取值范围**：不涉及。
	OsModelartsNodeCluster *string `json:"os.modelarts.node/cluster,omitempty"`

	// **参数解释**：节点绑定的逻辑池。 **取值范围**：不涉及。
	OsModelartsNodeElasticQuota *string `json:"os.modelarts.node/elastic.quota,omitempty"`

	// **参数解释**：节点所在的节点池id。 **取值范围**：不涉及。
	OsModelartsNodeNodepool *string `json:"os.modelarts.node/nodepool,omitempty"`

	// **参数解释**：批量创建批次标识。 **取值范围**：不涉及。
	OsModelartsNodeBatchUid *string `json:"os.modelarts.node/batch.uid,omitempty"`

	// **参数解释**：批量创建批次名称。 **取值范围**：不涉及。
	OsModelartsNodeBatchName *string `json:"os.modelarts.node/batch.name,omitempty"`

	// **参数解释**：批量创建批次类型。 **取值范围**：可选值如下：   - hyperinstance：超节点。
	OsModelartsNodeBatchType *string `json:"os.modelarts.node/batch.type,omitempty"`

	// **参数解释**：批量创建的节点个数。 **取值范围**：不涉及。
	OsModelartsNodeBatchCount *string `json:"os.modelarts.node/batch.count,omitempty"`

	// **参数解释**：节点的资源id。 **取值范围**：不涉及。
	OsModelartsResourceId *string `json:"os.modelarts/resource.id,omitempty"`

	// **参数解释**：节点的租户id，记录节点创建在哪个租户账号下。 **取值范围**：不涉及。
	OsModelartsTenantDomainId *string `json:"os.modelarts/tenant.domain.id,omitempty"`

	// **参数解释**：节点的项目id，记录节点创建在租户账号下哪个项目中。 **取值范围**：不涉及。
	OsModelartsTenantProjectId *string `json:"os.modelarts/tenant.project.id,omitempty"`

	// **参数解释**：节点计费状态。 **取值范围**：可选值如下： - 0：正常状态。 - 1：冻结状态。 - 2：删除状态或者终止状态。
	OsModelartsBillingStatus *string `json:"os.modelarts/billing.status,omitempty"`

	// **参数解释**：标识该节点是否被整柜作业独占。当被某个整柜作业独占时，该标签存在，标签的值为独占的训练作业ID。 **取值范围**：不涉及。
	OsModelartsNodeVolcanoSchedulerCabinetExclusive *string `json:"os.modelarts.node/volcano.scheduler.cabinet-exclusive,omitempty"`

	// **参数解释**：节点所在tor交换机ip。多个tor交换机ip之间以中划线-分隔。 **取值范围**：不涉及。
	CceKubectlKubernetesIoCabinet *string `json:"cce.kubectl.kubernetes.io/cabinet,omitempty"`

	// **参数解释**：节点底层资源的实例ID，如超节点的ECS实例ID。 **取值范围**：不涉及。
	OsModelartsNodeUnderlyingInstanceId *string `json:"os.modelarts.node/underlying.instance.id,omitempty"`

	// **参数解释**：节点是否启用高可用冗余。 **取值范围**：   - true：开启   - false：未开启
	OsModelartsNodeHaRedundantEnabled *string `json:"os.modelarts.node/ha.redundant.enabled,omitempty"`

	// **参数解释**：节点所在的节点池名称,最小长度为2，最大长度为50的小写字母、中划线-、数字组成，由小写字母开头，不能 以-，-default结尾。 **取值范围**：不涉及。
	OsModelartsNodeNodepoolname *string `json:"os.modelarts.node/nodepoolname,omitempty"`
}

func (o NodeLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLabels struct{}"
	}

	return strings.Join([]string{"NodeLabels", string(data)}, " ")
}
