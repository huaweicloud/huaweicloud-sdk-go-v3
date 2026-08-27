package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNotebooksRequest Request Object
type ListNotebooksRequest struct {

	// **参数解释**：实例类别。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - DEFAULT：CodeLab免费规格实例，每个用户最多只能创建一个。 - NOTEBOOK：计费规格实例。  **默认取值**：NOTEBOOK。
	Feature *ListNotebooksRequestFeature `json:"feature,omitempty"`

	// **参数解释**：每一页显示实例的数量。 **约束限制**：不涉及。 **取值范围**：[10,20,50]。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：实例名称，支持模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：长度限制128个字符，支持大小写字母、数字、中划线和下划线。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：专属资源池ID。 **约束限制**：不涉及。 **取值范围**：长度最长为64个字符，最短为4个字符，支持小写字母、数字、中划线，且必须是小写字母开头，小写字母或数字结尾。 **默认取值**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：实例归属的用户ID，在大账号/有admin权限场景下生效，值通常为当前登录用户ID。 **约束限制**：不涉及。 **取值范围**：长度为32位小写字母、数字。 **默认取值**：不涉及。
	Owner *string `json:"owner,omitempty"`

	// **参数解释**：实例排序方式。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - ASC：升序 - DESC：降序  **默认取值**：DESC。
	SortDir *ListNotebooksRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序的字段，多个字段使用(“,”)逗号分隔。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线、下划线和逗号。 **默认取值**：不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：实例状态。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - INIT：初始化 - CREATING：创建中 - STARTING：启动中 - STOPPING：停止中 - DELETING：删除中 - RUNNING：运行中 - STOPPED：已停止 - SNAPSHOTTING：快照中(保存镜像时的状态) - CREATE_FAILED：创建失败 - START_FAILED：启动失败 - DELETE_FAILED：删除失败 - ERROR：错误 - DELETED：已删除 - FROZEN：冻结  **默认取值**：不涉及。
	Status *ListNotebooksRequestStatus `json:"status,omitempty"`

	// **参数解释**：工作空间ID。获取方法请参见[[查询工作空间列表](ListWorkspace.xml)](tag:hc,hk)。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：实例的机器规格编码，支持模糊匹配查询。如下规格仅供参考，实际支持的规格以具体区域为准。 modelarts.vm.cpu.2u：Intel CPU通用规格，用于快速数据探索和实验。 modelarts.vm.cpu.8u：Intel CPU算力增强型，适用于密集计算场景下运算。 **约束限制**：不支持专属资源池的自定义规格查询。 **取值范围**：长度限制1-256字符，支持数字、大小写字母、小数点、下划线或中划线。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。镜像的ID可通过调用[[查询支持的镜像列表](https://support.huaweicloud.com/api-modelarts/ListImage.html)](tag:hc)[[查询支持的镜像列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListImage.html)](tag:hk)接口获取。 **约束限制**：不涉及。 **取值范围**：调用[[查询支持的镜像列表](https://support.huaweicloud.com/api-modelarts/ListImage.html)](tag:hc)[[查询支持的镜像列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListImage.html)](tag:hk)接口获取的合法镜像ID列表。 **默认取值**：不涉及。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)](tag:hc)[[查询Notebook实例列表接口](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListNotebooks.html#section0)](tag:hk)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：实例计费类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - COMPUTE：计算资源计费 - STORAGE：存储资源计费 - ALL：所有计费类型  **默认取值**：不涉及。
	Billing *string `json:"billing,omitempty"`

	// **参数解释**：实例标签信息。 **约束限制**：不涉及。 **取值范围**：不以逗号，竖划线开头，不以逗号结尾，不出现连续的竖划线和逗号，允许中文、西文、葡文等语言以及空格_.:/=+-@特殊字符，且字符间以逗号或者竖划线分割。例：tag_key1|tag_value1,tag_key2|tag_value2。 **默认取值**：不涉及。
	Tags *string `json:"tags,omitempty"`

	// **参数解释**：SWR镜像路径，该参数是针对返回参数NotebookResp中Image的swr_path属性进行模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：长度限制2048个字符，支持数字、大小写字母、下划线、中划线、点号、冒号和斜杠，0-2048个字符。 **默认取值**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：专属资源池名称，支持模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：长度限制1-64字符，支持数字、大小写字母和中划线。 **默认取值**：不涉及。
	PoolName *string `json:"pool_name,omitempty"`

	// **参数解释**：实例描述信息，支持模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：长度限制为512字符，不可包含特殊字符<>。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：节点IP。 **约束限制**：不涉及。 **取值范围**：正确的IPv4地址，暂不支持IPv6地址。 **默认取值**：不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**：实例创建用户名称，支持模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：长度限制1-256字符，支持数字、大小写字母、小数点、下划线或中划线。 **默认取值**：不涉及。
	Username *string `json:"username,omitempty"`
}

func (o ListNotebooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotebooksRequest struct{}"
	}

	return strings.Join([]string{"ListNotebooksRequest", string(data)}, " ")
}

type ListNotebooksRequestFeature struct {
	value string
}

type ListNotebooksRequestFeatureEnum struct {
	DEFAULT  ListNotebooksRequestFeature
	NOTEBOOK ListNotebooksRequestFeature
}

func GetListNotebooksRequestFeatureEnum() ListNotebooksRequestFeatureEnum {
	return ListNotebooksRequestFeatureEnum{
		DEFAULT: ListNotebooksRequestFeature{
			value: "DEFAULT",
		},
		NOTEBOOK: ListNotebooksRequestFeature{
			value: "NOTEBOOK",
		},
	}
}

func (c ListNotebooksRequestFeature) Value() string {
	return c.value
}

func (c ListNotebooksRequestFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotebooksRequestFeature) UnmarshalJSON(b []byte) error {
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

type ListNotebooksRequestSortDir struct {
	value string
}

type ListNotebooksRequestSortDirEnum struct {
	ASC  ListNotebooksRequestSortDir
	DESC ListNotebooksRequestSortDir
}

func GetListNotebooksRequestSortDirEnum() ListNotebooksRequestSortDirEnum {
	return ListNotebooksRequestSortDirEnum{
		ASC: ListNotebooksRequestSortDir{
			value: "ASC",
		},
		DESC: ListNotebooksRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListNotebooksRequestSortDir) Value() string {
	return c.value
}

func (c ListNotebooksRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotebooksRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListNotebooksRequestStatus struct {
	value string
}

type ListNotebooksRequestStatusEnum struct {
	CREATE_FAILED ListNotebooksRequestStatus
	CREATING      ListNotebooksRequestStatus
	DELETED       ListNotebooksRequestStatus
	DELETE_FAILED ListNotebooksRequestStatus
	DELETING      ListNotebooksRequestStatus
	ERROR         ListNotebooksRequestStatus
	FROZEN        ListNotebooksRequestStatus
	INIT          ListNotebooksRequestStatus
	RUNNING       ListNotebooksRequestStatus
	SNAPSHOTTING  ListNotebooksRequestStatus
	STARTING      ListNotebooksRequestStatus
	START_FAILED  ListNotebooksRequestStatus
	STOPPED       ListNotebooksRequestStatus
	STOPPING      ListNotebooksRequestStatus
}

func GetListNotebooksRequestStatusEnum() ListNotebooksRequestStatusEnum {
	return ListNotebooksRequestStatusEnum{
		CREATE_FAILED: ListNotebooksRequestStatus{
			value: "CREATE_FAILED",
		},
		CREATING: ListNotebooksRequestStatus{
			value: "CREATING",
		},
		DELETED: ListNotebooksRequestStatus{
			value: "DELETED",
		},
		DELETE_FAILED: ListNotebooksRequestStatus{
			value: "DELETE_FAILED",
		},
		DELETING: ListNotebooksRequestStatus{
			value: "DELETING",
		},
		ERROR: ListNotebooksRequestStatus{
			value: "ERROR",
		},
		FROZEN: ListNotebooksRequestStatus{
			value: "FROZEN",
		},
		INIT: ListNotebooksRequestStatus{
			value: "INIT",
		},
		RUNNING: ListNotebooksRequestStatus{
			value: "RUNNING",
		},
		SNAPSHOTTING: ListNotebooksRequestStatus{
			value: "SNAPSHOTTING",
		},
		STARTING: ListNotebooksRequestStatus{
			value: "STARTING",
		},
		START_FAILED: ListNotebooksRequestStatus{
			value: "START_FAILED",
		},
		STOPPED: ListNotebooksRequestStatus{
			value: "STOPPED",
		},
		STOPPING: ListNotebooksRequestStatus{
			value: "STOPPING",
		},
	}
}

func (c ListNotebooksRequestStatus) Value() string {
	return c.value
}

func (c ListNotebooksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotebooksRequestStatus) UnmarshalJSON(b []byte) error {
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
