package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMergeRequestCommentsByLineRequest Request Object
type ShowMergeRequestCommentsByLineRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid int32 `json:"merge_request_iid"`

	Line *int32 `json:"line,omitempty"`

	WithCommitComments *bool `json:"with_commit_comments,omitempty"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过100000。
	Path *string `json:"path,omitempty"`

	View *ShowMergeRequestCommentsByLineRequestView `json:"view,omitempty"`

	// 合并请求中原分支与目标分支的共同基准点
	BaseSha *string `json:"base_sha,omitempty"`

	// 合并请求中，从共同基准点开始到原分支方向的第一个提交点
	StartSha *string `json:"start_sha,omitempty"`

	// 合并请求中原分支指向的提交点
	HeadSha *string `json:"head_sha,omitempty"`
}

func (o ShowMergeRequestCommentsByLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestCommentsByLineRequest struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestCommentsByLineRequest", string(data)}, " ")
}

type ShowMergeRequestCommentsByLineRequestView struct {
	value string
}

type ShowMergeRequestCommentsByLineRequestViewEnum struct {
	BASIC ShowMergeRequestCommentsByLineRequestView
}

func GetShowMergeRequestCommentsByLineRequestViewEnum() ShowMergeRequestCommentsByLineRequestViewEnum {
	return ShowMergeRequestCommentsByLineRequestViewEnum{
		BASIC: ShowMergeRequestCommentsByLineRequestView{
			value: "basic",
		},
	}
}

func (c ShowMergeRequestCommentsByLineRequestView) Value() string {
	return c.value
}

func (c ShowMergeRequestCommentsByLineRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestCommentsByLineRequestView) UnmarshalJSON(b []byte) error {
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
