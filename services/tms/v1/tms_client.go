package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/tms/v1/model"
)

type TmsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewTmsClient(hcClient *httpclient.HcHttpClient) *TmsClient {
	return &TmsClient{HcClient: hcClient}
}

func TmsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ChangeAssociatedResourceOpenStatus 开通或关闭关联资源标签继承能力
//
// 开通或关闭关联资源标签继承能力
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ChangeAssociatedResourceOpenStatus(request *model.ChangeAssociatedResourceOpenStatusRequest) (*model.ChangeAssociatedResourceOpenStatusResponse, error) {
	requestDef := GenReqDefForChangeAssociatedResourceOpenStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAssociatedResourceOpenStatusResponse), nil
	}
}

// ChangeAssociatedResourceOpenStatusInvoker 开通或关闭关联资源标签继承能力
func (c *TmsClient) ChangeAssociatedResourceOpenStatusInvoker(request *model.ChangeAssociatedResourceOpenStatusRequest) *ChangeAssociatedResourceOpenStatusInvoker {
	requestDef := GenReqDefForChangeAssociatedResourceOpenStatus()
	return &ChangeAssociatedResourceOpenStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAssociatedResourceRules 批量启用规则
//
// 按照规则和region来批量启用规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) CreateAssociatedResourceRules(request *model.CreateAssociatedResourceRulesRequest) (*model.CreateAssociatedResourceRulesResponse, error) {
	requestDef := GenReqDefForCreateAssociatedResourceRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssociatedResourceRulesResponse), nil
	}
}

// CreateAssociatedResourceRulesInvoker 批量启用规则
func (c *TmsClient) CreateAssociatedResourceRulesInvoker(request *model.CreateAssociatedResourceRulesRequest) *CreateAssociatedResourceRulesInvoker {
	requestDef := GenReqDefForCreateAssociatedResourceRules()
	return &CreateAssociatedResourceRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePredefineTags 创建预定义标签
//
// 用于创建预定标签。用户创建预定义标签后，可以使用预定义标签来给资源创建标签。该接口支持幂等特性和处理批量数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) CreatePredefineTags(request *model.CreatePredefineTagsRequest) (*model.CreatePredefineTagsResponse, error) {
	requestDef := GenReqDefForCreatePredefineTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePredefineTagsResponse), nil
	}
}

// CreatePredefineTagsInvoker 创建预定义标签
func (c *TmsClient) CreatePredefineTagsInvoker(request *model.CreatePredefineTagsRequest) *CreatePredefineTagsInvoker {
	requestDef := GenReqDefForCreatePredefineTags()
	return &CreatePredefineTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceTag 批量添加标签
//
// 用于给云服务的多个资源添加标签，每个资源最多可添加10个标签，每次最多支持批量操作20个资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

// CreateResourceTagInvoker 批量添加标签
func (c *TmsClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAssociatedResourceRule 关闭规则
//
// 通过指定setting_name和region_id来关闭指定的规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) DeleteAssociatedResourceRule(request *model.DeleteAssociatedResourceRuleRequest) (*model.DeleteAssociatedResourceRuleResponse, error) {
	requestDef := GenReqDefForDeleteAssociatedResourceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssociatedResourceRuleResponse), nil
	}
}

// DeleteAssociatedResourceRuleInvoker 关闭规则
func (c *TmsClient) DeleteAssociatedResourceRuleInvoker(request *model.DeleteAssociatedResourceRuleRequest) *DeleteAssociatedResourceRuleInvoker {
	requestDef := GenReqDefForDeleteAssociatedResourceRule()
	return &DeleteAssociatedResourceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePredefineTags 删除预定义标签
//
// 用于删除预定标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) DeletePredefineTags(request *model.DeletePredefineTagsRequest) (*model.DeletePredefineTagsResponse, error) {
	requestDef := GenReqDefForDeletePredefineTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePredefineTagsResponse), nil
	}
}

// DeletePredefineTagsInvoker 删除预定义标签
func (c *TmsClient) DeletePredefineTagsInvoker(request *model.DeletePredefineTagsRequest) *DeletePredefineTagsInvoker {
	requestDef := GenReqDefForDeletePredefineTags()
	return &DeletePredefineTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 批量移除标签
//
// 用于批量移除云服务多个资源的标签，每个资源最多支持移除10个标签，每次最多支持批量操作20个资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 批量移除标签
func (c *TmsClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本列表
//
// 查询标签管理服务的API版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本列表
func (c *TmsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssociatedResourceRules 查询当前规则列表
//
// 查询当前规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListAssociatedResourceRules(request *model.ListAssociatedResourceRulesRequest) (*model.ListAssociatedResourceRulesResponse, error) {
	requestDef := GenReqDefForListAssociatedResourceRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssociatedResourceRulesResponse), nil
	}
}

// ListAssociatedResourceRulesInvoker 查询当前规则列表
func (c *TmsClient) ListAssociatedResourceRulesInvoker(request *model.ListAssociatedResourceRulesRequest) *ListAssociatedResourceRulesInvoker {
	requestDef := GenReqDefForListAssociatedResourceRules()
	return &ListAssociatedResourceRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssociatedResourceSettings 查询当前规则的配置列表
//
// 查询当前规则的配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListAssociatedResourceSettings(request *model.ListAssociatedResourceSettingsRequest) (*model.ListAssociatedResourceSettingsResponse, error) {
	requestDef := GenReqDefForListAssociatedResourceSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssociatedResourceSettingsResponse), nil
	}
}

// ListAssociatedResourceSettingsInvoker 查询当前规则的配置列表
func (c *TmsClient) ListAssociatedResourceSettingsInvoker(request *model.ListAssociatedResourceSettingsRequest) *ListAssociatedResourceSettingsInvoker {
	requestDef := GenReqDefForListAssociatedResourceSettings()
	return &ListAssociatedResourceSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPredefineTags 查询预定义标签列表
//
// 用于查询预定义标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListPredefineTags(request *model.ListPredefineTagsRequest) (*model.ListPredefineTagsResponse, error) {
	requestDef := GenReqDefForListPredefineTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPredefineTagsResponse), nil
	}
}

// ListPredefineTagsInvoker 查询预定义标签列表
func (c *TmsClient) ListPredefineTagsInvoker(request *model.ListPredefineTagsRequest) *ListPredefineTagsInvoker {
	requestDef := GenReqDefForListPredefineTags()
	return &ListPredefineTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProviders 查询标签管理支持的服务
//
// 查询标签管理支持的服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListProviders(request *model.ListProvidersRequest) (*model.ListProvidersResponse, error) {
	requestDef := GenReqDefForListProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvidersResponse), nil
	}
}

// ListProvidersInvoker 查询标签管理支持的服务
func (c *TmsClient) ListProvidersInvoker(request *model.ListProvidersRequest) *ListProvidersInvoker {
	requestDef := GenReqDefForListProviders()
	return &ListProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResource 根据标签过滤资源
//
// 根据标签过滤资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListResource(request *model.ListResourceRequest) (*model.ListResourceResponse, error) {
	requestDef := GenReqDefForListResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceResponse), nil
	}
}

// ListResourceInvoker 根据标签过滤资源
func (c *TmsClient) ListResourceInvoker(request *model.ListResourceRequest) *ListResourceInvoker {
	requestDef := GenReqDefForListResource()
	return &ListResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagKeys 查询标签键列表
//
// 查询指定区域的所有标签键.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListTagKeys(request *model.ListTagKeysRequest) (*model.ListTagKeysResponse, error) {
	requestDef := GenReqDefForListTagKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagKeysResponse), nil
	}
}

// ListTagKeysInvoker 查询标签键列表
func (c *TmsClient) ListTagKeysInvoker(request *model.ListTagKeysRequest) *ListTagKeysInvoker {
	requestDef := GenReqDefForListTagKeys()
	return &ListTagKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagValues 查询标签值列表
//
// 查询指定区域的标签键下的所有标签值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListTagValues(request *model.ListTagValuesRequest) (*model.ListTagValuesResponse, error) {
	requestDef := GenReqDefForListTagValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagValuesResponse), nil
	}
}

// ListTagValuesInvoker 查询标签值列表
func (c *TmsClient) ListTagValuesInvoker(request *model.ListTagValuesRequest) *ListTagValuesInvoker {
	requestDef := GenReqDefForListTagValues()
	return &ListTagValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询标签列表
//
// 查询指定区域和实例类型中租户的所有标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询标签列表
func (c *TmsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询API版本号详情
//
// 查询指定的标签管理服务API版本号详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询API版本号详情
func (c *TmsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssociatedResourceOpenStatus 查询当前用户的关联资源标签继承能力的开通状态
//
// 查询当前用户的关联资源标签继承能力的开通状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ShowAssociatedResourceOpenStatus(request *model.ShowAssociatedResourceOpenStatusRequest) (*model.ShowAssociatedResourceOpenStatusResponse, error) {
	requestDef := GenReqDefForShowAssociatedResourceOpenStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssociatedResourceOpenStatusResponse), nil
	}
}

// ShowAssociatedResourceOpenStatusInvoker 查询当前用户的关联资源标签继承能力的开通状态
func (c *TmsClient) ShowAssociatedResourceOpenStatusInvoker(request *model.ShowAssociatedResourceOpenStatusRequest) *ShowAssociatedResourceOpenStatusInvoker {
	requestDef := GenReqDefForShowAssociatedResourceOpenStatus()
	return &ShowAssociatedResourceOpenStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTag 查询资源标签
//
// 查询单个资源上的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

// ShowResourceTagInvoker 查询资源标签
func (c *TmsClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTagQuota 查询标签配额
//
// 查询标签的配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) ShowTagQuota(request *model.ShowTagQuotaRequest) (*model.ShowTagQuotaResponse, error) {
	requestDef := GenReqDefForShowTagQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagQuotaResponse), nil
	}
}

// ShowTagQuotaInvoker 查询标签配额
func (c *TmsClient) ShowTagQuotaInvoker(request *model.ShowTagQuotaRequest) *ShowTagQuotaInvoker {
	requestDef := GenReqDefForShowTagQuota()
	return &ShowTagQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssociatedResourceRules 更新规则
//
// 更新规则。规则更新的信息会覆盖原有内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) UpdateAssociatedResourceRules(request *model.UpdateAssociatedResourceRulesRequest) (*model.UpdateAssociatedResourceRulesResponse, error) {
	requestDef := GenReqDefForUpdateAssociatedResourceRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssociatedResourceRulesResponse), nil
	}
}

// UpdateAssociatedResourceRulesInvoker 更新规则
func (c *TmsClient) UpdateAssociatedResourceRulesInvoker(request *model.UpdateAssociatedResourceRulesRequest) *UpdateAssociatedResourceRulesInvoker {
	requestDef := GenReqDefForUpdateAssociatedResourceRules()
	return &UpdateAssociatedResourceRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePredefineTags 修改预定义标签
//
// 修改预定义标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *TmsClient) UpdatePredefineTags(request *model.UpdatePredefineTagsRequest) (*model.UpdatePredefineTagsResponse, error) {
	requestDef := GenReqDefForUpdatePredefineTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePredefineTagsResponse), nil
	}
}

// UpdatePredefineTagsInvoker 修改预定义标签
func (c *TmsClient) UpdatePredefineTagsInvoker(request *model.UpdatePredefineTagsRequest) *UpdatePredefineTagsInvoker {
	requestDef := GenReqDefForUpdatePredefineTags()
	return &UpdatePredefineTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
