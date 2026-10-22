package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotGroupDetail struct {

	// **参数解释** 快照一致性组的ID。 **取值范围** 不涉及。
	Id string `json:"id"`

	// **参数解释** 快照一致性组的项目ID。 **取值范围** 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释** 快照一致性组的创建时间。 **取值范围** 不涉及。
	CreatedAt string `json:"created_at"`

	// **参数解释** 快照一致性组的状态。 **取值范围** - creating：创建中。 - available：创建成功。 - error：创建失败。 - deleting: 删除中。 - deleted：删除成功。 - error_deleting：删除失败。
	Status string `json:"status"`

	// **参数解释** 快照一致性组的更新时间。 **取值范围** 不涉及。
	UpdatedAt string `json:"updated_at"`

	// **参数解释** 快照一致性组的名称。 **取值范围** 不涉及。
	Name string `json:"name"`

	// **参数解释** 快照一致性组的描述。 **取值范围** 不涉及。
	Description string `json:"description"`

	// **参数解释** 快照一致性组的企业项目id。 **取值范围** 不涉及。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释** 快照一致性组的标签。 标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。  标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。 **取值范围** 不涉及。
	Tags map[string]string `json:"tags"`

	// **参数解释** 云服务器id。 **取值范围** 不涉及。
	ServerId string `json:"server_id"`
}

func (o SnapshotGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotGroupDetail struct{}"
	}

	return strings.Join([]string{"SnapshotGroupDetail", string(data)}, " ")
}
