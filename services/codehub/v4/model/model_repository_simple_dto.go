package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositorySimpleDto struct {

	// **参数解释：** 仓库ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库描述信息。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 仓库名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 仓库完整名称。
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 仓库路径。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 仓库完整路径。
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 是否归档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 仓库ssh地址。
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// **参数解释：** 仓库http地址。
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// **参数解释：** 仓库页面链接。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释：** 仓库readme文件链接。
	ReadmeUrl *string `json:"readme_url,omitempty"`

	// **参数解释：** 仓库所属项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 仓库所属项目名称。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 仓库开发模式。 **取值范围：** - normal - CR
	DevelopMode *RepositorySimpleDtoDevelopMode `json:"develop_mode,omitempty"`

	// **参数解释：** 审核状态。
	ModerationResult *bool `json:"moderation_result,omitempty"`
}

func (o RepositorySimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositorySimpleDto struct{}"
	}

	return strings.Join([]string{"RepositorySimpleDto", string(data)}, " ")
}

type RepositorySimpleDtoDevelopMode struct {
	value string
}

type RepositorySimpleDtoDevelopModeEnum struct {
	NORMAL RepositorySimpleDtoDevelopMode
	CR     RepositorySimpleDtoDevelopMode
}

func GetRepositorySimpleDtoDevelopModeEnum() RepositorySimpleDtoDevelopModeEnum {
	return RepositorySimpleDtoDevelopModeEnum{
		NORMAL: RepositorySimpleDtoDevelopMode{
			value: "normal",
		},
		CR: RepositorySimpleDtoDevelopMode{
			value: "CR",
		},
	}
}

func (c RepositorySimpleDtoDevelopMode) Value() string {
	return c.value
}

func (c RepositorySimpleDtoDevelopMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositorySimpleDtoDevelopMode) UnmarshalJSON(b []byte) error {
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
