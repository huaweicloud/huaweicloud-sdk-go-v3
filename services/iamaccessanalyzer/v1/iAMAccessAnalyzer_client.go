package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iamaccessanalyzer/v1/model"
)

type IAMAccessAnalyzerClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIAMAccessAnalyzerClient(hcClient *httpclient.HcHttpClient) *IAMAccessAnalyzerClient {
	return &IAMAccessAnalyzerClient{HcClient: hcClient}
}

func IAMAccessAnalyzerClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAnalyzer 创建分析器
//
// 为您的账号或者组织创建分析器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) CreateAnalyzer(request *model.CreateAnalyzerRequest) (*model.CreateAnalyzerResponse, error) {
	requestDef := GenReqDefForCreateAnalyzer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnalyzerResponse), nil
	}
}

// CreateAnalyzerInvoker 创建分析器
func (c *IAMAccessAnalyzerClient) CreateAnalyzerInvoker(request *model.CreateAnalyzerRequest) *CreateAnalyzerInvoker {
	requestDef := GenReqDefForCreateAnalyzer()
	return &CreateAnalyzerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAnalyzer 删除指定的分析器
//
// 删除指定的分析器。分析器生成的所有检查结果都将被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) DeleteAnalyzer(request *model.DeleteAnalyzerRequest) (*model.DeleteAnalyzerResponse, error) {
	requestDef := GenReqDefForDeleteAnalyzer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAnalyzerResponse), nil
	}
}

// DeleteAnalyzerInvoker 删除指定的分析器
func (c *IAMAccessAnalyzerClient) DeleteAnalyzerInvoker(request *model.DeleteAnalyzerRequest) *DeleteAnalyzerInvoker {
	requestDef := GenReqDefForDeleteAnalyzer()
	return &DeleteAnalyzerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAnalyzers 检索分析器的列表
//
// 检索分析器的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ListAnalyzers(request *model.ListAnalyzersRequest) (*model.ListAnalyzersResponse, error) {
	requestDef := GenReqDefForListAnalyzers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAnalyzersResponse), nil
	}
}

// ListAnalyzersInvoker 检索分析器的列表
func (c *IAMAccessAnalyzerClient) ListAnalyzersInvoker(request *model.ListAnalyzersRequest) *ListAnalyzersInvoker {
	requestDef := GenReqDefForListAnalyzers()
	return &ListAnalyzersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAnalyzer 显示指定的分析器
//
// 检索有关指定分析器的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ShowAnalyzer(request *model.ShowAnalyzerRequest) (*model.ShowAnalyzerResponse, error) {
	requestDef := GenReqDefForShowAnalyzer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAnalyzerResponse), nil
	}
}

// ShowAnalyzerInvoker 显示指定的分析器
func (c *IAMAccessAnalyzerClient) ShowAnalyzerInvoker(request *model.ShowAnalyzerRequest) *ShowAnalyzerInvoker {
	requestDef := GenReqDefForShowAnalyzer()
	return &ShowAnalyzerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartResourceScan 立即开始扫描应用于指定资源的策略
//
// 立即开始扫描应用于指定资源的策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) StartResourceScan(request *model.StartResourceScanRequest) (*model.StartResourceScanResponse, error) {
	requestDef := GenReqDefForStartResourceScan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResourceScanResponse), nil
	}
}

// StartResourceScanInvoker 立即开始扫描应用于指定资源的策略
func (c *IAMAccessAnalyzerClient) StartResourceScanInvoker(request *model.StartResourceScanRequest) *StartResourceScanInvoker {
	requestDef := GenReqDefForStartResourceScan()
	return &StartResourceScanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyArchiveRule 应用存档规则
//
// 以追溯方式将存档规则应用于符合存档规则条件的现有结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ApplyArchiveRule(request *model.ApplyArchiveRuleRequest) (*model.ApplyArchiveRuleResponse, error) {
	requestDef := GenReqDefForApplyArchiveRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyArchiveRuleResponse), nil
	}
}

// ApplyArchiveRuleInvoker 应用存档规则
func (c *IAMAccessAnalyzerClient) ApplyArchiveRuleInvoker(request *model.ApplyArchiveRuleRequest) *ApplyArchiveRuleInvoker {
	requestDef := GenReqDefForApplyArchiveRule()
	return &ApplyArchiveRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateArchiveRule 为指定的分析器创建存档规则
//
// 为指定的分析器创建存档规则。存档规则会自动存档符合您在创建规则时所定义条件的新结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) CreateArchiveRule(request *model.CreateArchiveRuleRequest) (*model.CreateArchiveRuleResponse, error) {
	requestDef := GenReqDefForCreateArchiveRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateArchiveRuleResponse), nil
	}
}

// CreateArchiveRuleInvoker 为指定的分析器创建存档规则
func (c *IAMAccessAnalyzerClient) CreateArchiveRuleInvoker(request *model.CreateArchiveRuleRequest) *CreateArchiveRuleInvoker {
	requestDef := GenReqDefForCreateArchiveRule()
	return &CreateArchiveRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteArchiveRule 删除指定的存档规则
//
// 删除指定的存档规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) DeleteArchiveRule(request *model.DeleteArchiveRuleRequest) (*model.DeleteArchiveRuleResponse, error) {
	requestDef := GenReqDefForDeleteArchiveRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteArchiveRuleResponse), nil
	}
}

// DeleteArchiveRuleInvoker 删除指定的存档规则
func (c *IAMAccessAnalyzerClient) DeleteArchiveRuleInvoker(request *model.DeleteArchiveRuleRequest) *DeleteArchiveRuleInvoker {
	requestDef := GenReqDefForDeleteArchiveRule()
	return &DeleteArchiveRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArchiveRules 检索为指定分析器创建的存档规则的列表
//
// 检索为指定分析器创建的存档规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ListArchiveRules(request *model.ListArchiveRulesRequest) (*model.ListArchiveRulesResponse, error) {
	requestDef := GenReqDefForListArchiveRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArchiveRulesResponse), nil
	}
}

// ListArchiveRulesInvoker 检索为指定分析器创建的存档规则的列表
func (c *IAMAccessAnalyzerClient) ListArchiveRulesInvoker(request *model.ListArchiveRulesRequest) *ListArchiveRulesInvoker {
	requestDef := GenReqDefForListArchiveRules()
	return &ListArchiveRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowArchiveRule 检索有关存档规则的信息
//
// 检索有关存档规则的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ShowArchiveRule(request *model.ShowArchiveRuleRequest) (*model.ShowArchiveRuleResponse, error) {
	requestDef := GenReqDefForShowArchiveRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowArchiveRuleResponse), nil
	}
}

// ShowArchiveRuleInvoker 检索有关存档规则的信息
func (c *IAMAccessAnalyzerClient) ShowArchiveRuleInvoker(request *model.ShowArchiveRuleRequest) *ShowArchiveRuleInvoker {
	requestDef := GenReqDefForShowArchiveRule()
	return &ShowArchiveRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateArchiveRule 更新指定存档规则的条件和值
//
// 更新指定存档规则的条件和值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) UpdateArchiveRule(request *model.UpdateArchiveRuleRequest) (*model.UpdateArchiveRuleResponse, error) {
	requestDef := GenReqDefForUpdateArchiveRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateArchiveRuleResponse), nil
	}
}

// UpdateArchiveRuleInvoker 更新指定存档规则的条件和值
func (c *IAMAccessAnalyzerClient) UpdateArchiveRuleInvoker(request *model.UpdateArchiveRuleRequest) *UpdateArchiveRuleInvoker {
	requestDef := GenReqDefForUpdateArchiveRule()
	return &UpdateArchiveRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFindings 检索指定分析器生成的访问分析结果列表
//
// 检索指定分析器生成的访问分析结果列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ListFindings(request *model.ListFindingsRequest) (*model.ListFindingsResponse, error) {
	requestDef := GenReqDefForListFindings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFindingsResponse), nil
	}
}

// ListFindingsInvoker 检索指定分析器生成的访问分析结果列表
func (c *IAMAccessAnalyzerClient) ListFindingsInvoker(request *model.ListFindingsRequest) *ListFindingsInvoker {
	requestDef := GenReqDefForListFindings()
	return &ListFindingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFinding 检索有关指定结果的信息
//
// 检索有关指定结果的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ShowFinding(request *model.ShowFindingRequest) (*model.ShowFindingResponse, error) {
	requestDef := GenReqDefForShowFinding()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFindingResponse), nil
	}
}

// ShowFindingInvoker 检索有关指定结果的信息
func (c *IAMAccessAnalyzerClient) ShowFindingInvoker(request *model.ShowFindingRequest) *ShowFindingInvoker {
	requestDef := GenReqDefForShowFinding()
	return &ShowFindingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFindings 更新指定结果的状态
//
// 更新指定访问分析结果的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) UpdateFindings(request *model.UpdateFindingsRequest) (*model.UpdateFindingsResponse, error) {
	requestDef := GenReqDefForUpdateFindings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFindingsResponse), nil
	}
}

// UpdateFindingsInvoker 更新指定结果的状态
func (c *IAMAccessAnalyzerClient) UpdateFindingsInvoker(request *model.UpdateFindingsRequest) *UpdateFindingsInvoker {
	requestDef := GenReqDefForUpdateFindings()
	return &UpdateFindingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessPreview 创建访问预览
//
// 创建访问预览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) CreateAccessPreview(request *model.CreateAccessPreviewRequest) (*model.CreateAccessPreviewResponse, error) {
	requestDef := GenReqDefForCreateAccessPreview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessPreviewResponse), nil
	}
}

// CreateAccessPreviewInvoker 创建访问预览
func (c *IAMAccessAnalyzerClient) CreateAccessPreviewInvoker(request *model.CreateAccessPreviewRequest) *CreateAccessPreviewInvoker {
	requestDef := GenReqDefForCreateAccessPreview()
	return &CreateAccessPreviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPreviewFindings 获取相关预览生成的分析结果
//
// 获取相关预览生成的分析结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ListAccessPreviewFindings(request *model.ListAccessPreviewFindingsRequest) (*model.ListAccessPreviewFindingsResponse, error) {
	requestDef := GenReqDefForListAccessPreviewFindings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPreviewFindingsResponse), nil
	}
}

// ListAccessPreviewFindingsInvoker 获取相关预览生成的分析结果
func (c *IAMAccessAnalyzerClient) ListAccessPreviewFindingsInvoker(request *model.ListAccessPreviewFindingsRequest) *ListAccessPreviewFindingsInvoker {
	requestDef := GenReqDefForListAccessPreviewFindings()
	return &ListAccessPreviewFindingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPreviews 获取所有访问预览
//
// 获取所有访问预览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ListAccessPreviews(request *model.ListAccessPreviewsRequest) (*model.ListAccessPreviewsResponse, error) {
	requestDef := GenReqDefForListAccessPreviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPreviewsResponse), nil
	}
}

// ListAccessPreviewsInvoker 获取所有访问预览
func (c *IAMAccessAnalyzerClient) ListAccessPreviewsInvoker(request *model.ListAccessPreviewsRequest) *ListAccessPreviewsInvoker {
	requestDef := GenReqDefForListAccessPreviews()
	return &ListAccessPreviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessPreview 获取相关访问预览的信息
//
// 获取相关访问预览的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ShowAccessPreview(request *model.ShowAccessPreviewRequest) (*model.ShowAccessPreviewResponse, error) {
	requestDef := GenReqDefForShowAccessPreview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessPreviewResponse), nil
	}
}

// ShowAccessPreviewInvoker 获取相关访问预览的信息
func (c *IAMAccessAnalyzerClient) ShowAccessPreviewInvoker(request *model.ShowAccessPreviewRequest) *ShowAccessPreviewInvoker {
	requestDef := GenReqDefForShowAccessPreview()
	return &ShowAccessPreviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagResource 向指定资源添加标签
//
// 向指定资源添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) TagResource(request *model.TagResourceRequest) (*model.TagResourceResponse, error) {
	requestDef := GenReqDefForTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagResourceResponse), nil
	}
}

// TagResourceInvoker 向指定资源添加标签
func (c *IAMAccessAnalyzerClient) TagResourceInvoker(request *model.TagResourceRequest) *TagResourceInvoker {
	requestDef := GenReqDefForTagResource()
	return &TagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UntagResource 从指定资源中删除标签
//
// 从指定资源中删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) UntagResource(request *model.UntagResourceRequest) (*model.UntagResourceResponse, error) {
	requestDef := GenReqDefForUntagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagResourceResponse), nil
	}
}

// UntagResourceInvoker 从指定资源中删除标签
func (c *IAMAccessAnalyzerClient) UntagResourceInvoker(request *model.UntagResourceRequest) *UntagResourceInvoker {
	requestDef := GenReqDefForUntagResource()
	return &UntagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckNoNewAccess 校验策略是否有新访问权限
//
// 校验策略是否有新访问权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) CheckNoNewAccess(request *model.CheckNoNewAccessRequest) (*model.CheckNoNewAccessResponse, error) {
	requestDef := GenReqDefForCheckNoNewAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNoNewAccessResponse), nil
	}
}

// CheckNoNewAccessInvoker 校验策略是否有新访问权限
func (c *IAMAccessAnalyzerClient) CheckNoNewAccessInvoker(request *model.CheckNoNewAccessRequest) *CheckNoNewAccessInvoker {
	requestDef := GenReqDefForCheckNoNewAccess()
	return &CheckNoNewAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidatePolicy 校验策略
//
// 校验策略并返回结果列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IAMAccessAnalyzerClient) ValidatePolicy(request *model.ValidatePolicyRequest) (*model.ValidatePolicyResponse, error) {
	requestDef := GenReqDefForValidatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidatePolicyResponse), nil
	}
}

// ValidatePolicyInvoker 校验策略
func (c *IAMAccessAnalyzerClient) ValidatePolicyInvoker(request *model.ValidatePolicyRequest) *ValidatePolicyInvoker {
	requestDef := GenReqDefForValidatePolicy()
	return &ValidatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
