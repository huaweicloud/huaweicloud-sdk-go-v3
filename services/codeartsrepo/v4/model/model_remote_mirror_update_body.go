package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RemoteMirrorUpdateBody struct {

	// **参数解释：** 源仓地址。 **取值范围：** 不涉及。 **约束限制：** 以 http:// 或 https:// 开头。 **默认取值：** 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释：**  分支同步。 **取值范围：** - all，同步全部分支。 - default，同步默认分支。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	SyncBranchType *RemoteMirrorUpdateBodySyncBranchType `json:"sync_branch_type,omitempty"`

	// **参数解释：**  开启定时同步 **取值范围：** - true，不开启定时同步。 - false，开启定时同步。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	MirroringEnabled *bool `json:"mirroring_enabled,omitempty"`

	// **参数解释：**  拓展点UUID。 **取值范围：** 不涉及。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`
}

func (o RemoteMirrorUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteMirrorUpdateBody struct{}"
	}

	return strings.Join([]string{"RemoteMirrorUpdateBody", string(data)}, " ")
}

type RemoteMirrorUpdateBodySyncBranchType struct {
	value string
}

type RemoteMirrorUpdateBodySyncBranchTypeEnum struct {
	ALL     RemoteMirrorUpdateBodySyncBranchType
	DEFAULT RemoteMirrorUpdateBodySyncBranchType
}

func GetRemoteMirrorUpdateBodySyncBranchTypeEnum() RemoteMirrorUpdateBodySyncBranchTypeEnum {
	return RemoteMirrorUpdateBodySyncBranchTypeEnum{
		ALL: RemoteMirrorUpdateBodySyncBranchType{
			value: "all",
		},
		DEFAULT: RemoteMirrorUpdateBodySyncBranchType{
			value: "default",
		},
	}
}

func (c RemoteMirrorUpdateBodySyncBranchType) Value() string {
	return c.value
}

func (c RemoteMirrorUpdateBodySyncBranchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemoteMirrorUpdateBodySyncBranchType) UnmarshalJSON(b []byte) error {
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
