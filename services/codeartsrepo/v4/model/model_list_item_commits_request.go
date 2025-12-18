package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListItemCommitsRequest Request Object
type ListItemCommitsRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 工作项Id。 **约束限制：** 不涉及  **取值范围：** 字符串长度不少于1，不超过128。
	ItemId string `json:"item_id"`

	// **参数解释：** 工作项关联的提交类型。 **约束限制：** 不涉及  **取值范围：** - commit，提交。 - branch，分支。 - mergerequest，合并请求。
	Type *ListItemCommitsRequestType `json:"type,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListItemCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListItemCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListItemCommitsRequest", string(data)}, " ")
}

type ListItemCommitsRequestType struct {
	value string
}

type ListItemCommitsRequestTypeEnum struct {
	COMMIT       ListItemCommitsRequestType
	BRANCH       ListItemCommitsRequestType
	MERGEREQUEST ListItemCommitsRequestType
}

func GetListItemCommitsRequestTypeEnum() ListItemCommitsRequestTypeEnum {
	return ListItemCommitsRequestTypeEnum{
		COMMIT: ListItemCommitsRequestType{
			value: "commit",
		},
		BRANCH: ListItemCommitsRequestType{
			value: "branch",
		},
		MERGEREQUEST: ListItemCommitsRequestType{
			value: "mergerequest",
		},
	}
}

func (c ListItemCommitsRequestType) Value() string {
	return c.value
}

func (c ListItemCommitsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListItemCommitsRequestType) UnmarshalJSON(b []byte) error {
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
