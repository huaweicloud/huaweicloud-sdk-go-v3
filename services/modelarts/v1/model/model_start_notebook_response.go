package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartNotebookResponse Response Object
type StartNotebookResponse struct {

	// **参数解释**：实例初始化进度。
	ActionProgress *[]JobProgress `json:"action_progress,omitempty"`

	// **参数解释**：实例描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：本地IDE（如PyCharm、VS Code）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
	Endpoints *[]EndpointsRes `json:"endpoints,omitempty"`

	// **参数解释**：实例失败原因。 **取值范围**：不涉及。
	FailReason *string `json:"fail_reason,omitempty"`

	// **参数解释**：实例规格， 1.当用户选择系统规格时，返回值为系统规格码； 2.当用户创建实例时选择了自定义规格，则此字段会固定返回\"custom.flavor.spec.code\"。 **取值范围**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	CustomSpec *NotebookCustomSpecRep `json:"custom_spec,omitempty"`

	// **参数解释**：实例ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	Image *Image `json:"image,omitempty"`

	Lease *Lease `json:"lease,omitempty"`

	// **参数解释**：实例名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	Pool *Pool `json:"pool,omitempty"`

	// **参数解释**：实例状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化 - CREATING：创建中 - STARTING：启动中 - STOPPING：停止中 - DELETING：删除中 - RUNNING：运行中 - STOPPED：已停止 - SNAPSHOTTING：快照中(保存镜像时的状态) - CREATE_FAILED：创建失败 - START_FAILED：启动失败 - DELETE_FAILED：删除失败 - ERROR：错误 - DELETED：已删除 - FROZEN：冻结
	Status *StartNotebookResponseStatus `json:"status,omitempty"`

	// **参数解释**：Notebook鉴权使用的token信息。 **取值范围**：不涉及。
	Token *string `json:"token,omitempty"`

	// **参数解释**：实例访问的URL。 **取值范围**：不涉及。
	Url *string `json:"url,omitempty"`

	Volume *VolumeRes `json:"volume,omitempty"`

	// **参数解释**：工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：实例类别。 **取值范围**：枚举类型，取值如下： - DEFAULT：CodeLab免费规格实例，每个用户最多只能创建一个。 - NOTEBOOK：计费规格实例。
	Feature *string `json:"feature,omitempty"`

	// **参数解释**：计费资源类型。枚举类型，取值如下： - STORAGE：存储资源计费。 - COMPUTE：计算资源计费。 - ALL：所有计费类型。
	BillingItems *[]StartNotebookResponseBillingItems `json:"billing_items,omitempty"`

	User *UserResponse `json:"user,omitempty"`

	Affinity *AffinityType `json:"affinity,omitempty"`

	RunUser *RunUserInfo `json:"run_user,omitempty"`

	// **参数解释**：扩展存储信息
	DataVolumes *[]VolumeResponse `json:"data_volumes,omitempty"`

	// **参数解释**：实例所在节点ip。 **取值范围**：不涉及。
	Ip *string `json:"ip,omitempty"`

	UserVpc *UserVpcResponse `json:"user_vpc,omitempty"`

	// **参数解释**：用户ID。 **取值范围**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：是否需要默认创建用户secret，默认为true。 **取值范围**：不涉及。
	IsNeedCredentials *bool `json:"is_need_credentials,omitempty"`

	// **参数解释**：jupyter version版本号。 **取值范围**：不涉及。
	JupyterVersion *string `json:"jupyter_version,omitempty"`

	// **参数解释**：实例标签。
	Tags           *[]TmsTagResponse `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o StartNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNotebookResponse struct{}"
	}

	return strings.Join([]string{"StartNotebookResponse", string(data)}, " ")
}

type StartNotebookResponseStatus struct {
	value string
}

type StartNotebookResponseStatusEnum struct {
	CREATE_FAILED StartNotebookResponseStatus
	CREATING      StartNotebookResponseStatus
	DELETED       StartNotebookResponseStatus
	DELETE_FAILED StartNotebookResponseStatus
	DELETING      StartNotebookResponseStatus
	ERROR         StartNotebookResponseStatus
	FROZEN        StartNotebookResponseStatus
	INIT          StartNotebookResponseStatus
	RUNNING       StartNotebookResponseStatus
	SNAPSHOTTING  StartNotebookResponseStatus
	STARTING      StartNotebookResponseStatus
	START_FAILED  StartNotebookResponseStatus
	STOPPED       StartNotebookResponseStatus
	STOPPING      StartNotebookResponseStatus
}

func GetStartNotebookResponseStatusEnum() StartNotebookResponseStatusEnum {
	return StartNotebookResponseStatusEnum{
		CREATE_FAILED: StartNotebookResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: StartNotebookResponseStatus{
			value: "CREATING",
		},
		DELETED: StartNotebookResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: StartNotebookResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: StartNotebookResponseStatus{
			value: "DELETING",
		},
		ERROR: StartNotebookResponseStatus{
			value: "ERROR",
		},
		FROZEN: StartNotebookResponseStatus{
			value: "FROZEN",
		},
		INIT: StartNotebookResponseStatus{
			value: "INIT",
		},
		RUNNING: StartNotebookResponseStatus{
			value: "RUNNING",
		},
		SNAPSHOTTING: StartNotebookResponseStatus{
			value: "SNAPSHOTTING",
		},
		STARTING: StartNotebookResponseStatus{
			value: "STARTING",
		},
		START_FAILED: StartNotebookResponseStatus{
			value: "START_FAILED",
		},
		STOPPED: StartNotebookResponseStatus{
			value: "STOPPED",
		},
		STOPPING: StartNotebookResponseStatus{
			value: "STOPPING",
		},
	}
}

func (c StartNotebookResponseStatus) Value() string {
	return c.value
}

func (c StartNotebookResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartNotebookResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type StartNotebookResponseBillingItems struct {
	value string
}

type StartNotebookResponseBillingItemsEnum struct {
	STORAGE StartNotebookResponseBillingItems
	COMPUTE StartNotebookResponseBillingItems
	ALL     StartNotebookResponseBillingItems
}

func GetStartNotebookResponseBillingItemsEnum() StartNotebookResponseBillingItemsEnum {
	return StartNotebookResponseBillingItemsEnum{
		STORAGE: StartNotebookResponseBillingItems{
			value: "STORAGE",
		},
		COMPUTE: StartNotebookResponseBillingItems{
			value: "COMPUTE",
		},
		ALL: StartNotebookResponseBillingItems{
			value: "ALL",
		},
	}
}

func (c StartNotebookResponseBillingItems) Value() string {
	return c.value
}

func (c StartNotebookResponseBillingItems) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartNotebookResponseBillingItems) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
