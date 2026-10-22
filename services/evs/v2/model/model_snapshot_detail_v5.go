package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotDetailV5 struct {

	// **参数解释** 云硬盘快照ID。 **取值范围** 不涉及。
	Id string `json:"id"`

	// **参数解释** 云硬盘快照名称。 名称以autobk_snapshot_为前缀的快照是创建云硬盘备份时系统自动创建的，请不要进行“删除云硬盘快照”和“回滚快照到云硬盘”操作。 **取值范围** 不涉及。
	Name string `json:"name"`

	// **参数解释** 云硬盘快照描述信息。 **取值范围** 不涉及。
	Description string `json:"description"`

	// **参数解释** 云硬盘快照创建时间。 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX **取值范围** 不涉及。
	CreatedAt string `json:"created_at"`

	// **参数解释** 云硬盘快照更新时间。 **取值范围** 不涉及。
	UpdatedAt string `json:"updated_at"`

	// **参数解释** 快照所属的云硬盘ID。 **取值范围** 不涉及。
	VolumeId string `json:"volume_id"`

	// **参数解释** 云硬盘快照大小，单位为GiB。 **取值范围** 不涉及。
	Size int32 `json:"size"`

	// **参数解释** 云硬盘快照的状态。 **取值范围** 具体请参见[云硬盘快照状态](evs_04_0041.xml)。
	Status string `json:"status"`

	// **参数解释** 项目ID。 **取值范围** 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释** 快照所属的企业项目ID。 **取值范围** 不涉及。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释** 是否为加密快照。 **取值范围** 不涉及。
	Encrypted bool `json:"encrypted"`

	// **参数解释** 快照所属的云硬盘ID的密钥ID。 **取值范围** 不涉及。
	CmkId string `json:"cmk_id"`

	// **参数解释** 快照类型。 **取值范围** 枚举值：standard/flash。
	Category string `json:"category"`

	// **参数解释** 快照所属AZ。 **取值范围** 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释** 快照所属标签。 标签的key：String类型，最大长度36个字符。 由英文字母、数字、下划线、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。  标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。 **取值范围** 不涉及。
	Tags map[string]string `json:"tags"`

	// **参数解释** 是否开启了快照极速可用功能。 **取值范围** true：开启。仅SSD系列云盘支持开启该功能。 false：关闭。即快照为未开启快照极速可用功能的存量快照。
	InstantAccess bool `json:"instant_access"`

	// **参数解释** 快照保留的时间。 **取值范围** 不涉及。
	RetentionAt string `json:"retention_at"`

	// **参数解释** 设置快照极速可用功能的保留时间，保留时间到期后自动关闭快照极速可用功能。 **取值范围** 默认与参数retention_at的值一致。
	InstantAccessRetentionAt string `json:"instant_access_retention_at"`

	// **参数解释** 是否增量快照。 **取值范围** true，false。
	Incremental bool `json:"incremental"`

	// **参数解释** 快照创建来源。 **取值范围** auto：自动创建快照。 user：手动创建快照。 copy：复制快照。
	SnapshotType string `json:"snapshot_type"`

	// **参数解释** 快照创建进度，单位为百分比。 **取值范围** 不涉及。
	Progress string `json:"progress"`

	// **参数解释** 加密快照的加密算法。 **取值范围** 不涉及。
	EncryptAlgorithm string `json:"encrypt_algorithm"`

	// **参数解释** 快照所属的快照链信息。 **取值范围** 不涉及。
	SnapshotChains []SnapshotChains `json:"snapshot_chains"`

	// **参数解释** 快照所属的一致性快照组id。 **取值范围** 不涉及。
	SnapshotGroupId *string `json:"snapshot_group_id,omitempty"`
}

func (o SnapshotDetailV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotDetailV5 struct{}"
	}

	return strings.Join([]string{"SnapshotDetailV5", string(data)}, " ")
}
