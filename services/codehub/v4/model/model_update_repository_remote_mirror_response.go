package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRepositoryRemoteMirrorResponse Response Object
type UpdateRepositoryRemoteMirrorResponse struct {

	// **参数解释：**  源仓地址。 **约束限制：** 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释：**  开启定时同步 **取值范围：** - true，不开启定时同步。 - false，开启定时同步。 **约束限制：** 不涉及。
	MirroringEnabled *bool `json:"mirroring_enabled,omitempty"`

	// **参数解释：**  拓展点UUID。 **约束限制：** 不涉及。
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`

	// **参数解释：**  分支同步。 **取值范围：** - all，同步全部分支。 - default，同步默认分支。 **约束限制：** 不涉及。
	SyncBranchType *UpdateRepositoryRemoteMirrorResponseSyncBranchType `json:"sync_branch_type,omitempty"`
	HttpStatusCode int                                                 `json:"-"`
}

func (o UpdateRepositoryRemoteMirrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryRemoteMirrorResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryRemoteMirrorResponse", string(data)}, " ")
}

type UpdateRepositoryRemoteMirrorResponseSyncBranchType struct {
	value string
}

type UpdateRepositoryRemoteMirrorResponseSyncBranchTypeEnum struct {
	ALL     UpdateRepositoryRemoteMirrorResponseSyncBranchType
	DEFAULT UpdateRepositoryRemoteMirrorResponseSyncBranchType
}

func GetUpdateRepositoryRemoteMirrorResponseSyncBranchTypeEnum() UpdateRepositoryRemoteMirrorResponseSyncBranchTypeEnum {
	return UpdateRepositoryRemoteMirrorResponseSyncBranchTypeEnum{
		ALL: UpdateRepositoryRemoteMirrorResponseSyncBranchType{
			value: "all",
		},
		DEFAULT: UpdateRepositoryRemoteMirrorResponseSyncBranchType{
			value: "default",
		},
	}
}

func (c UpdateRepositoryRemoteMirrorResponseSyncBranchType) Value() string {
	return c.value
}

func (c UpdateRepositoryRemoteMirrorResponseSyncBranchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRepositoryRemoteMirrorResponseSyncBranchType) UnmarshalJSON(b []byte) error {
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
