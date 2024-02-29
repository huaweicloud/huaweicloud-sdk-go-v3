package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idmeclassicapi/v1/model"
)

type IDMEClassicAPIClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIDMEClassicAPIClient(hcClient *httpclient.HcHttpClient) *IDMEClassicAPIClient {
	return &IDMEClassicAPIClient{HcClient: hcClient}
}

func IDMEClassicAPIClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowBatchCheckinUsingPost XDM_批量检入VersionModel
//
// 根据主对象ID批量检入版本对象，小版本升版。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchCheckinUsingPost(request *model.ShowBatchCheckinUsingPostRequest) (*model.ShowBatchCheckinUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchCheckinUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchCheckinUsingPostResponse), nil
	}
}

// ShowBatchCheckinUsingPostInvoker XDM_批量检入VersionModel
func (c *IDMEClassicAPIClient) ShowBatchCheckinUsingPostInvoker(request *model.ShowBatchCheckinUsingPostRequest) *ShowBatchCheckinUsingPostInvoker {
	requestDef := GenReqDefForShowBatchCheckinUsingPost()
	return &ShowBatchCheckinUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchCheckoutAndUpdateUsingPost XDM_批量检出并更新VersionModel
//
// 根据主对象ID批量检出对象并根据传入字段批量更新版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchCheckoutAndUpdateUsingPost(request *model.ShowBatchCheckoutAndUpdateUsingPostRequest) (*model.ShowBatchCheckoutAndUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchCheckoutAndUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchCheckoutAndUpdateUsingPostResponse), nil
	}
}

// ShowBatchCheckoutAndUpdateUsingPostInvoker XDM_批量检出并更新VersionModel
func (c *IDMEClassicAPIClient) ShowBatchCheckoutAndUpdateUsingPostInvoker(request *model.ShowBatchCheckoutAndUpdateUsingPostRequest) *ShowBatchCheckoutAndUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowBatchCheckoutAndUpdateUsingPost()
	return &ShowBatchCheckoutAndUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchCheckoutUsingPost XDM_批量检出VersionModel
//
// 根据主对象ID批量检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchCheckoutUsingPost(request *model.ShowBatchCheckoutUsingPostRequest) (*model.ShowBatchCheckoutUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchCheckoutUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchCheckoutUsingPostResponse), nil
	}
}

// ShowBatchCheckoutUsingPostInvoker XDM_批量检出VersionModel
func (c *IDMEClassicAPIClient) ShowBatchCheckoutUsingPostInvoker(request *model.ShowBatchCheckoutUsingPostRequest) *ShowBatchCheckoutUsingPostInvoker {
	requestDef := GenReqDefForShowBatchCheckoutUsingPost()
	return &ShowBatchCheckoutUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchCreateUsingPost 批量创建实例
//
// 批量创建指定数据模型的数据实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchCreateUsingPost(request *model.ShowBatchCreateUsingPostRequest) (*model.ShowBatchCreateUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchCreateUsingPostResponse), nil
	}
}

// ShowBatchCreateUsingPostInvoker 批量创建实例
func (c *IDMEClassicAPIClient) ShowBatchCreateUsingPostInvoker(request *model.ShowBatchCreateUsingPostRequest) *ShowBatchCreateUsingPostInvoker {
	requestDef := GenReqDefForShowBatchCreateUsingPost()
	return &ShowBatchCreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchDeleteBranchUsingPost XDM_批量删除VersionModel最新分支版本下所有小版本
//
// 根据主对象ID&amp;业务版本列表，批量删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchDeleteBranchUsingPost(request *model.ShowBatchDeleteBranchUsingPostRequest) (*model.ShowBatchDeleteBranchUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchDeleteBranchUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchDeleteBranchUsingPostResponse), nil
	}
}

// ShowBatchDeleteBranchUsingPostInvoker XDM_批量删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) ShowBatchDeleteBranchUsingPostInvoker(request *model.ShowBatchDeleteBranchUsingPostRequest) *ShowBatchDeleteBranchUsingPostInvoker {
	requestDef := GenReqDefForShowBatchDeleteBranchUsingPost()
	return &ShowBatchDeleteBranchUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchDeleteUsingPost 批量删除实例
//
// 根据数据实例的唯一编码，批量删除指定数据模型中的多个数据实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchDeleteUsingPost(request *model.ShowBatchDeleteUsingPostRequest) (*model.ShowBatchDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchDeleteUsingPostResponse), nil
	}
}

// ShowBatchDeleteUsingPostInvoker 批量删除实例
func (c *IDMEClassicAPIClient) ShowBatchDeleteUsingPostInvoker(request *model.ShowBatchDeleteUsingPostRequest) *ShowBatchDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowBatchDeleteUsingPost()
	return &ShowBatchDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchGetUsingPost 批量查询实例
//
// 根据多个数据实例的唯一编码，批量查询实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchGetUsingPost(request *model.ShowBatchGetUsingPostRequest) (*model.ShowBatchGetUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchGetUsingPostResponse), nil
	}
}

// ShowBatchGetUsingPostInvoker 批量查询实例
func (c *IDMEClassicAPIClient) ShowBatchGetUsingPostInvoker(request *model.ShowBatchGetUsingPostRequest) *ShowBatchGetUsingPostInvoker {
	requestDef := GenReqDefForShowBatchGetUsingPost()
	return &ShowBatchGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchLogicalDeleteBranchUsingPost XDM_批量软删除VersionModel最新分支版本下所有小版本
//
// 批量软删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchLogicalDeleteBranchUsingPost(request *model.ShowBatchLogicalDeleteBranchUsingPostRequest) (*model.ShowBatchLogicalDeleteBranchUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchLogicalDeleteBranchUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchLogicalDeleteBranchUsingPostResponse), nil
	}
}

// ShowBatchLogicalDeleteBranchUsingPostInvoker XDM_批量软删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) ShowBatchLogicalDeleteBranchUsingPostInvoker(request *model.ShowBatchLogicalDeleteBranchUsingPostRequest) *ShowBatchLogicalDeleteBranchUsingPostInvoker {
	requestDef := GenReqDefForShowBatchLogicalDeleteBranchUsingPost()
	return &ShowBatchLogicalDeleteBranchUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchLogicalDeleteUsingPost 批量软删除实例
//
// 根据数据实例的唯一编码，批量软删除指定数据模型中的多个数据实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchLogicalDeleteUsingPost(request *model.ShowBatchLogicalDeleteUsingPostRequest) (*model.ShowBatchLogicalDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchLogicalDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchLogicalDeleteUsingPostResponse), nil
	}
}

// ShowBatchLogicalDeleteUsingPostInvoker 批量软删除实例
func (c *IDMEClassicAPIClient) ShowBatchLogicalDeleteUsingPostInvoker(request *model.ShowBatchLogicalDeleteUsingPostRequest) *ShowBatchLogicalDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowBatchLogicalDeleteUsingPost()
	return &ShowBatchLogicalDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchReviseAndUpdateUsingPost XDM_批量修订且更新VersionModel。
//
// 根据主对象ID批量修订对象并根据传入字段更新主对象+版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchReviseAndUpdateUsingPost(request *model.ShowBatchReviseAndUpdateUsingPostRequest) (*model.ShowBatchReviseAndUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchReviseAndUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchReviseAndUpdateUsingPostResponse), nil
	}
}

// ShowBatchReviseAndUpdateUsingPostInvoker XDM_批量修订且更新VersionModel。
func (c *IDMEClassicAPIClient) ShowBatchReviseAndUpdateUsingPostInvoker(request *model.ShowBatchReviseAndUpdateUsingPostRequest) *ShowBatchReviseAndUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowBatchReviseAndUpdateUsingPost()
	return &ShowBatchReviseAndUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchReviseUsingPost XDM_批量修订VersionModel。
//
// 根据主对象ID批量修订对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchReviseUsingPost(request *model.ShowBatchReviseUsingPostRequest) (*model.ShowBatchReviseUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchReviseUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchReviseUsingPostResponse), nil
	}
}

// ShowBatchReviseUsingPostInvoker XDM_批量修订VersionModel。
func (c *IDMEClassicAPIClient) ShowBatchReviseUsingPostInvoker(request *model.ShowBatchReviseUsingPostRequest) *ShowBatchReviseUsingPostInvoker {
	requestDef := GenReqDefForShowBatchReviseUsingPost()
	return &ShowBatchReviseUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUndoCheckoutByAdminUsingPost XDM_管理员批量撤销检出VersionModel
//
// 管理员根据主对象ID批量撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUndoCheckoutByAdminUsingPost(request *model.ShowBatchUndoCheckoutByAdminUsingPostRequest) (*model.ShowBatchUndoCheckoutByAdminUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUndoCheckoutByAdminUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUndoCheckoutByAdminUsingPostResponse), nil
	}
}

// ShowBatchUndoCheckoutByAdminUsingPostInvoker XDM_管理员批量撤销检出VersionModel
func (c *IDMEClassicAPIClient) ShowBatchUndoCheckoutByAdminUsingPostInvoker(request *model.ShowBatchUndoCheckoutByAdminUsingPostRequest) *ShowBatchUndoCheckoutByAdminUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUndoCheckoutByAdminUsingPost()
	return &ShowBatchUndoCheckoutByAdminUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUndoCheckoutUsingPost XDM_批量撤销检出VersionModel
//
// 根据主对象ID批量撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUndoCheckoutUsingPost(request *model.ShowBatchUndoCheckoutUsingPostRequest) (*model.ShowBatchUndoCheckoutUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUndoCheckoutUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUndoCheckoutUsingPostResponse), nil
	}
}

// ShowBatchUndoCheckoutUsingPostInvoker XDM_批量撤销检出VersionModel
func (c *IDMEClassicAPIClient) ShowBatchUndoCheckoutUsingPostInvoker(request *model.ShowBatchUndoCheckoutUsingPostRequest) *ShowBatchUndoCheckoutUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUndoCheckoutUsingPost()
	return &ShowBatchUndoCheckoutUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUpdateAndCheckinUsingPost XDM_批量更新并检入VersionModel
//
// 根据传入字段批量更新版本对象并根据主对象ID批量检入对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUpdateAndCheckinUsingPost(request *model.ShowBatchUpdateAndCheckinUsingPostRequest) (*model.ShowBatchUpdateAndCheckinUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUpdateAndCheckinUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUpdateAndCheckinUsingPostResponse), nil
	}
}

// ShowBatchUpdateAndCheckinUsingPostInvoker XDM_批量更新并检入VersionModel
func (c *IDMEClassicAPIClient) ShowBatchUpdateAndCheckinUsingPostInvoker(request *model.ShowBatchUpdateAndCheckinUsingPostRequest) *ShowBatchUpdateAndCheckinUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUpdateAndCheckinUsingPost()
	return &ShowBatchUpdateAndCheckinUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUpdateByAdminUsingPost XDM_管理员批量更新VersionModel指定属性
//
// 以管理员身份批量更新指定版本实例上的基础信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUpdateByAdminUsingPost(request *model.ShowBatchUpdateByAdminUsingPostRequest) (*model.ShowBatchUpdateByAdminUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUpdateByAdminUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUpdateByAdminUsingPostResponse), nil
	}
}

// ShowBatchUpdateByAdminUsingPostInvoker XDM_管理员批量更新VersionModel指定属性
func (c *IDMEClassicAPIClient) ShowBatchUpdateByAdminUsingPostInvoker(request *model.ShowBatchUpdateByAdminUsingPostRequest) *ShowBatchUpdateByAdminUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUpdateByAdminUsingPost()
	return &ShowBatchUpdateByAdminUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUpdateUsingPost 批量更新实例
//
// 批量更新指定数据模型中的多个实例数据。如果某个实例的唯一编码不存在，该实例不做任何更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUpdateUsingPost(request *model.ShowBatchUpdateUsingPostRequest) (*model.ShowBatchUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUpdateUsingPostResponse), nil
	}
}

// ShowBatchUpdateUsingPostInvoker 批量更新实例
func (c *IDMEClassicAPIClient) ShowBatchUpdateUsingPostInvoker(request *model.ShowBatchUpdateUsingPostRequest) *ShowBatchUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUpdateUsingPost()
	return &ShowBatchUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUpdateVersionUsingPost XDM_批量升版最新版本对象VersionModel的版本号
//
// 根据ID批量升版最新版本对象数据的版本号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowBatchUpdateVersionUsingPost(request *model.ShowBatchUpdateVersionUsingPostRequest) (*model.ShowBatchUpdateVersionUsingPostResponse, error) {
	requestDef := GenReqDefForShowBatchUpdateVersionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUpdateVersionUsingPostResponse), nil
	}
}

// ShowBatchUpdateVersionUsingPostInvoker XDM_批量升版最新版本对象VersionModel的版本号
func (c *IDMEClassicAPIClient) ShowBatchUpdateVersionUsingPostInvoker(request *model.ShowBatchUpdateVersionUsingPostRequest) *ShowBatchUpdateVersionUsingPostInvoker {
	requestDef := GenReqDefForShowBatchUpdateVersionUsingPost()
	return &ShowBatchUpdateVersionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckinUsingPost XDM_检入VersionModel
//
// 根据主对象ID检入版本对象，按照设置的规则生成新的业务版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCheckinUsingPost(request *model.ShowCheckinUsingPostRequest) (*model.ShowCheckinUsingPostResponse, error) {
	requestDef := GenReqDefForShowCheckinUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckinUsingPostResponse), nil
	}
}

// ShowCheckinUsingPostInvoker XDM_检入VersionModel
func (c *IDMEClassicAPIClient) ShowCheckinUsingPostInvoker(request *model.ShowCheckinUsingPostRequest) *ShowCheckinUsingPostInvoker {
	requestDef := GenReqDefForShowCheckinUsingPost()
	return &ShowCheckinUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckoutAndUpdateUsingPost XDM_检出并更新VersionModel
//
// 根据主对象ID检出对象并根据传入字段更新版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCheckoutAndUpdateUsingPost(request *model.ShowCheckoutAndUpdateUsingPostRequest) (*model.ShowCheckoutAndUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowCheckoutAndUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckoutAndUpdateUsingPostResponse), nil
	}
}

// ShowCheckoutAndUpdateUsingPostInvoker XDM_检出并更新VersionModel
func (c *IDMEClassicAPIClient) ShowCheckoutAndUpdateUsingPostInvoker(request *model.ShowCheckoutAndUpdateUsingPostRequest) *ShowCheckoutAndUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowCheckoutAndUpdateUsingPost()
	return &ShowCheckoutAndUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckoutUsingPost XDM_检出VersionModel
//
// 根据主对象ID检出版本对象，复制生成一条新的版本记录且状态为已检出。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCheckoutUsingPost(request *model.ShowCheckoutUsingPostRequest) (*model.ShowCheckoutUsingPostResponse, error) {
	requestDef := GenReqDefForShowCheckoutUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckoutUsingPostResponse), nil
	}
}

// ShowCheckoutUsingPostInvoker XDM_检出VersionModel
func (c *IDMEClassicAPIClient) ShowCheckoutUsingPostInvoker(request *model.ShowCheckoutUsingPostRequest) *ShowCheckoutUsingPostInvoker {
	requestDef := GenReqDefForShowCheckoutUsingPost()
	return &ShowCheckoutUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCompareBusinessVersionUsingPost XDM_VersionModel业务版本对比
//
// 根据主对象id，输入版本号（大版本+小版本）进行版本属性与关系对比。（建议使用新的接口instance-attrs-comparison和\\ instance-relation-comparison比较属性和关系）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCompareBusinessVersionUsingPost(request *model.ShowCompareBusinessVersionUsingPostRequest) (*model.ShowCompareBusinessVersionUsingPostResponse, error) {
	requestDef := GenReqDefForShowCompareBusinessVersionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCompareBusinessVersionUsingPostResponse), nil
	}
}

// ShowCompareBusinessVersionUsingPostInvoker XDM_VersionModel业务版本对比
func (c *IDMEClassicAPIClient) ShowCompareBusinessVersionUsingPostInvoker(request *model.ShowCompareBusinessVersionUsingPostRequest) *ShowCompareBusinessVersionUsingPostInvoker {
	requestDef := GenReqDefForShowCompareBusinessVersionUsingPost()
	return &ShowCompareBusinessVersionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCountUsingPost 统计指定数据模型的实例总数
//
// 根据指定的查询条件，统计指定数据模型中的实例总数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCountUsingPost(request *model.ShowCountUsingPostRequest) (*model.ShowCountUsingPostResponse, error) {
	requestDef := GenReqDefForShowCountUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCountUsingPostResponse), nil
	}
}

// ShowCountUsingPostInvoker 统计指定数据模型的实例总数
func (c *IDMEClassicAPIClient) ShowCountUsingPostInvoker(request *model.ShowCountUsingPostRequest) *ShowCountUsingPostInvoker {
	requestDef := GenReqDefForShowCountUsingPost()
	return &ShowCountUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCreateUsingPost 创建实例
//
// 创建指定数据模型的数据实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowCreateUsingPost(request *model.ShowCreateUsingPostRequest) (*model.ShowCreateUsingPostResponse, error) {
	requestDef := GenReqDefForShowCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCreateUsingPostResponse), nil
	}
}

// ShowCreateUsingPostInvoker 创建实例
func (c *IDMEClassicAPIClient) ShowCreateUsingPostInvoker(request *model.ShowCreateUsingPostRequest) *ShowCreateUsingPostInvoker {
	requestDef := GenReqDefForShowCreateUsingPost()
	return &ShowCreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeleteBranchUsingPost XDM_删除VersionModel最新分支版本下所有小版本
//
// 根据masterid&amp;version删除最新大版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowDeleteBranchUsingPost(request *model.ShowDeleteBranchUsingPostRequest) (*model.ShowDeleteBranchUsingPostResponse, error) {
	requestDef := GenReqDefForShowDeleteBranchUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeleteBranchUsingPostResponse), nil
	}
}

// ShowDeleteBranchUsingPostInvoker XDM_删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) ShowDeleteBranchUsingPostInvoker(request *model.ShowDeleteBranchUsingPostRequest) *ShowDeleteBranchUsingPostInvoker {
	requestDef := GenReqDefForShowDeleteBranchUsingPost()
	return &ShowDeleteBranchUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeleteByConditionUsingPost 根据指定条件删除实例
//
// 通过此接口，删除指定条件查询返回的实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowDeleteByConditionUsingPost(request *model.ShowDeleteByConditionUsingPostRequest) (*model.ShowDeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForShowDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeleteByConditionUsingPostResponse), nil
	}
}

// ShowDeleteByConditionUsingPostInvoker 根据指定条件删除实例
func (c *IDMEClassicAPIClient) ShowDeleteByConditionUsingPostInvoker(request *model.ShowDeleteByConditionUsingPostRequest) *ShowDeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForShowDeleteByConditionUsingPost()
	return &ShowDeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeleteLatestVersionUsingPost XDM_删除VersionModel最新分支的最新版本
//
// 根据主对象ID入参，删除最新分支的最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowDeleteLatestVersionUsingPost(request *model.ShowDeleteLatestVersionUsingPostRequest) (*model.ShowDeleteLatestVersionUsingPostResponse, error) {
	requestDef := GenReqDefForShowDeleteLatestVersionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeleteLatestVersionUsingPostResponse), nil
	}
}

// ShowDeleteLatestVersionUsingPostInvoker XDM_删除VersionModel最新分支的最新版本
func (c *IDMEClassicAPIClient) ShowDeleteLatestVersionUsingPostInvoker(request *model.ShowDeleteLatestVersionUsingPostRequest) *ShowDeleteLatestVersionUsingPostInvoker {
	requestDef := GenReqDefForShowDeleteLatestVersionUsingPost()
	return &ShowDeleteLatestVersionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeleteUsingPost 删除实例
//
// 根据数据实例的唯一编码，删除指定数据模型中的一个数据实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowDeleteUsingPost(request *model.ShowDeleteUsingPostRequest) (*model.ShowDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeleteUsingPostResponse), nil
	}
}

// ShowDeleteUsingPostInvoker 删除实例
func (c *IDMEClassicAPIClient) ShowDeleteUsingPostInvoker(request *model.ShowDeleteUsingPostRequest) *ShowDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowDeleteUsingPost()
	return &ShowDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFindUsingPost 分页查询实例
//
// 分页查询指定数据模型中的所有实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowFindUsingPost(request *model.ShowFindUsingPostRequest) (*model.ShowFindUsingPostResponse, error) {
	requestDef := GenReqDefForShowFindUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFindUsingPostResponse), nil
	}
}

// ShowFindUsingPostInvoker 分页查询实例
func (c *IDMEClassicAPIClient) ShowFindUsingPostInvoker(request *model.ShowFindUsingPostRequest) *ShowFindUsingPostInvoker {
	requestDef := GenReqDefForShowFindUsingPost()
	return &ShowFindUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetAllVersionsUsingPost XDM_获取VersionModel版本列表
//
// 根据主对象ID，获取全量版本以及对应版本对象list属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetAllVersionsUsingPost(request *model.ShowGetAllVersionsUsingPostRequest) (*model.ShowGetAllVersionsUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetAllVersionsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetAllVersionsUsingPostResponse), nil
	}
}

// ShowGetAllVersionsUsingPostInvoker XDM_获取VersionModel版本列表
func (c *IDMEClassicAPIClient) ShowGetAllVersionsUsingPostInvoker(request *model.ShowGetAllVersionsUsingPostRequest) *ShowGetAllVersionsUsingPostInvoker {
	requestDef := GenReqDefForShowGetAllVersionsUsingPost()
	return &ShowGetAllVersionsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetByUniqueKeyUsingPost 根据唯一键为“是”的属性查询实例
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性查询实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetByUniqueKeyUsingPost(request *model.ShowGetByUniqueKeyUsingPostRequest) (*model.ShowGetByUniqueKeyUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetByUniqueKeyUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetByUniqueKeyUsingPostResponse), nil
	}
}

// ShowGetByUniqueKeyUsingPostInvoker 根据唯一键为“是”的属性查询实例
func (c *IDMEClassicAPIClient) ShowGetByUniqueKeyUsingPostInvoker(request *model.ShowGetByUniqueKeyUsingPostRequest) *ShowGetByUniqueKeyUsingPostInvoker {
	requestDef := GenReqDefForShowGetByUniqueKeyUsingPost()
	return &ShowGetByUniqueKeyUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetUsingPost 查询实例
//
// 根据一个数据实例的唯一编码，查询实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetUsingPost(request *model.ShowGetUsingPostRequest) (*model.ShowGetUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetUsingPostResponse), nil
	}
}

// ShowGetUsingPostInvoker 查询实例
func (c *IDMEClassicAPIClient) ShowGetUsingPostInvoker(request *model.ShowGetUsingPostRequest) *ShowGetUsingPostInvoker {
	requestDef := GenReqDefForShowGetUsingPost()
	return &ShowGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetVersionByMasterUsingPost XDM_获取VersionModel对应版本信息
//
// 根据Masterid和版本号和小版本号，返回对应版本属性，小版本号为空则返回最新小版本属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetVersionByMasterUsingPost(request *model.ShowGetVersionByMasterUsingPostRequest) (*model.ShowGetVersionByMasterUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetVersionByMasterUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetVersionByMasterUsingPostResponse), nil
	}
}

// ShowGetVersionByMasterUsingPostInvoker XDM_获取VersionModel对应版本信息
func (c *IDMEClassicAPIClient) ShowGetVersionByMasterUsingPostInvoker(request *model.ShowGetVersionByMasterUsingPostRequest) *ShowGetVersionByMasterUsingPostInvoker {
	requestDef := GenReqDefForShowGetVersionByMasterUsingPost()
	return &ShowGetVersionByMasterUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListUsingPost 查询实例的基础属性
//
// 根据查询条件分页返回模型基本属性信息且不级联查询（不支持扩展属性作为查询条件）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowListUsingPost(request *model.ShowListUsingPostRequest) (*model.ShowListUsingPostResponse, error) {
	requestDef := GenReqDefForShowListUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListUsingPostResponse), nil
	}
}

// ShowListUsingPostInvoker 查询实例的基础属性
func (c *IDMEClassicAPIClient) ShowListUsingPostInvoker(request *model.ShowListUsingPostRequest) *ShowListUsingPostInvoker {
	requestDef := GenReqDefForShowListUsingPost()
	return &ShowListUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteBranchUsingPost XDM_软删除VersionModel最新分支版本下所有小版本
//
// 软删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteBranchUsingPost(request *model.ShowLogicalDeleteBranchUsingPostRequest) (*model.ShowLogicalDeleteBranchUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteBranchUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteBranchUsingPostResponse), nil
	}
}

// ShowLogicalDeleteBranchUsingPostInvoker XDM_软删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) ShowLogicalDeleteBranchUsingPostInvoker(request *model.ShowLogicalDeleteBranchUsingPostRequest) *ShowLogicalDeleteBranchUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteBranchUsingPost()
	return &ShowLogicalDeleteBranchUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteByConditionUsingPost 根据指定条件软删除实例
//
// 通过此接口，软删除指定条件查询返回的实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPost(request *model.ShowLogicalDeleteByConditionUsingPostRequest) (*model.ShowLogicalDeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteByConditionUsingPostResponse), nil
	}
}

// ShowLogicalDeleteByConditionUsingPostInvoker 根据指定条件软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPostInvoker(request *model.ShowLogicalDeleteByConditionUsingPostRequest) *ShowLogicalDeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()
	return &ShowLogicalDeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteLatestVersionUsingPost XDM_软删除VersionModel最新分支的最新版本
//
// 根据主对象ID入参，软删最新分支的最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteLatestVersionUsingPost(request *model.ShowLogicalDeleteLatestVersionUsingPostRequest) (*model.ShowLogicalDeleteLatestVersionUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteLatestVersionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteLatestVersionUsingPostResponse), nil
	}
}

// ShowLogicalDeleteLatestVersionUsingPostInvoker XDM_软删除VersionModel最新分支的最新版本
func (c *IDMEClassicAPIClient) ShowLogicalDeleteLatestVersionUsingPostInvoker(request *model.ShowLogicalDeleteLatestVersionUsingPostRequest) *ShowLogicalDeleteLatestVersionUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteLatestVersionUsingPost()
	return &ShowLogicalDeleteLatestVersionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteUsingPost 软删除实例
//
// 根据数据实例的唯一编码，软删除指定数据模型中的一个数据实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPost(request *model.ShowLogicalDeleteUsingPostRequest) (*model.ShowLogicalDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteUsingPostResponse), nil
	}
}

// ShowLogicalDeleteUsingPostInvoker 软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPostInvoker(request *model.ShowLogicalDeleteUsingPostRequest) *ShowLogicalDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()
	return &ShowLogicalDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQueryUsingPost 根据“列表属性”为“是”的属性查询实例
//
// 当数据模型中存在“列表属性”为“是”的属性时，可通过此接口查询数据模型中的实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowQueryUsingPost(request *model.ShowQueryUsingPostRequest) (*model.ShowQueryUsingPostResponse, error) {
	requestDef := GenReqDefForShowQueryUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueryUsingPostResponse), nil
	}
}

// ShowQueryUsingPostInvoker 根据“列表属性”为“是”的属性查询实例
func (c *IDMEClassicAPIClient) ShowQueryUsingPostInvoker(request *model.ShowQueryUsingPostRequest) *ShowQueryUsingPostInvoker {
	requestDef := GenReqDefForShowQueryUsingPost()
	return &ShowQueryUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReviseAndUpdateUsingPost XDM_修订且更新VersionModel。
//
// 根据主对象ID修订对象并根据传入字段更新主对象+版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowReviseAndUpdateUsingPost(request *model.ShowReviseAndUpdateUsingPostRequest) (*model.ShowReviseAndUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowReviseAndUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReviseAndUpdateUsingPostResponse), nil
	}
}

// ShowReviseAndUpdateUsingPostInvoker XDM_修订且更新VersionModel。
func (c *IDMEClassicAPIClient) ShowReviseAndUpdateUsingPostInvoker(request *model.ShowReviseAndUpdateUsingPostRequest) *ShowReviseAndUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowReviseAndUpdateUsingPost()
	return &ShowReviseAndUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReviseUsingPost XDM_修订VersionModel
//
// 根据主对象ID修订对象，按照设置的规则生成新的业务版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowReviseUsingPost(request *model.ShowReviseUsingPostRequest) (*model.ShowReviseUsingPostResponse, error) {
	requestDef := GenReqDefForShowReviseUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReviseUsingPostResponse), nil
	}
}

// ShowReviseUsingPostInvoker XDM_修订VersionModel
func (c *IDMEClassicAPIClient) ShowReviseUsingPostInvoker(request *model.ShowReviseUsingPostRequest) *ShowReviseUsingPostInvoker {
	requestDef := GenReqDefForShowReviseUsingPost()
	return &ShowReviseUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSaveAllUsingPost 根据唯一键为“是”的属性更新实例数据
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性的英文名称更新该数据模型中实例的所有字段数据。如果更新的实例不存在，系统将自动创建该实例数据。
//
// 调用此接口时，建议传入该实例的所有字段信息。如果未传入某个字段，该字段的数据将更新为空值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowSaveAllUsingPost(request *model.ShowSaveAllUsingPostRequest) (*model.ShowSaveAllUsingPostResponse, error) {
	requestDef := GenReqDefForShowSaveAllUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSaveAllUsingPostResponse), nil
	}
}

// ShowSaveAllUsingPostInvoker 根据唯一键为“是”的属性更新实例数据
func (c *IDMEClassicAPIClient) ShowSaveAllUsingPostInvoker(request *model.ShowSaveAllUsingPostRequest) *ShowSaveAllUsingPostInvoker {
	requestDef := GenReqDefForShowSaveAllUsingPost()
	return &ShowSaveAllUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSaveAsUsingPost 另存版本对象的实例数据
//
// 版本对象的另存为接口（saveAs）用于创建一条与原版本对象实例数据相同的数据实例。该实例数据会完全复制原实例现有的数据，包括与其关联的主对象和分支对象，且新实例数据的版本号从初始值开始计算。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowSaveAsUsingPost(request *model.ShowSaveAsUsingPostRequest) (*model.ShowSaveAsUsingPostResponse, error) {
	requestDef := GenReqDefForShowSaveAsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSaveAsUsingPostResponse), nil
	}
}

// ShowSaveAsUsingPostInvoker 另存版本对象的实例数据
func (c *IDMEClassicAPIClient) ShowSaveAsUsingPostInvoker(request *model.ShowSaveAsUsingPostRequest) *ShowSaveAsUsingPostInvoker {
	requestDef := GenReqDefForShowSaveAsUsingPost()
	return &ShowSaveAsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSaveUsingPost 根据唯一键为“是”的属性更新实例的指定字段
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性的英文名称更新该数据模型中实例的指定字段数据。
//
// 如果更新的实例不存在，系统将自动创建该实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowSaveUsingPost(request *model.ShowSaveUsingPostRequest) (*model.ShowSaveUsingPostResponse, error) {
	requestDef := GenReqDefForShowSaveUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSaveUsingPostResponse), nil
	}
}

// ShowSaveUsingPostInvoker 根据唯一键为“是”的属性更新实例的指定字段
func (c *IDMEClassicAPIClient) ShowSaveUsingPostInvoker(request *model.ShowSaveUsingPostRequest) *ShowSaveUsingPostInvoker {
	requestDef := GenReqDefForShowSaveUsingPost()
	return &ShowSaveUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSelectUsingPost 查询实例的指定属性
//
// 根据查询条件及指定属性分页返回（不支持扩展属性作为选定属性列)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowSelectUsingPost(request *model.ShowSelectUsingPostRequest) (*model.ShowSelectUsingPostResponse, error) {
	requestDef := GenReqDefForShowSelectUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSelectUsingPostResponse), nil
	}
}

// ShowSelectUsingPostInvoker 查询实例的指定属性
func (c *IDMEClassicAPIClient) ShowSelectUsingPostInvoker(request *model.ShowSelectUsingPostRequest) *ShowSelectUsingPostInvoker {
	requestDef := GenReqDefForShowSelectUsingPost()
	return &ShowSelectUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStaticsUsingPost 查询指定数据模型的实例统计信息
//
// 根据指定函数，统计指定数据模型的实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowStaticsUsingPost(request *model.ShowStaticsUsingPostRequest) (*model.ShowStaticsUsingPostResponse, error) {
	requestDef := GenReqDefForShowStaticsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStaticsUsingPostResponse), nil
	}
}

// ShowStaticsUsingPostInvoker 查询指定数据模型的实例统计信息
func (c *IDMEClassicAPIClient) ShowStaticsUsingPostInvoker(request *model.ShowStaticsUsingPostRequest) *ShowStaticsUsingPostInvoker {
	requestDef := GenReqDefForShowStaticsUsingPost()
	return &ShowStaticsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUndoCheckoutByAdminUsingPost XDM_管理员撤销检出VersionModel
//
// 管理员根据主对象ID撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUndoCheckoutByAdminUsingPost(request *model.ShowUndoCheckoutByAdminUsingPostRequest) (*model.ShowUndoCheckoutByAdminUsingPostResponse, error) {
	requestDef := GenReqDefForShowUndoCheckoutByAdminUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUndoCheckoutByAdminUsingPostResponse), nil
	}
}

// ShowUndoCheckoutByAdminUsingPostInvoker XDM_管理员撤销检出VersionModel
func (c *IDMEClassicAPIClient) ShowUndoCheckoutByAdminUsingPostInvoker(request *model.ShowUndoCheckoutByAdminUsingPostRequest) *ShowUndoCheckoutByAdminUsingPostInvoker {
	requestDef := GenReqDefForShowUndoCheckoutByAdminUsingPost()
	return &ShowUndoCheckoutByAdminUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUndoCheckoutUsingPost XDM_撤销检出VersionModel
//
// 根据主对象ID撤销检出版本对象，删除新的版本记录且状态为已检入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUndoCheckoutUsingPost(request *model.ShowUndoCheckoutUsingPostRequest) (*model.ShowUndoCheckoutUsingPostResponse, error) {
	requestDef := GenReqDefForShowUndoCheckoutUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUndoCheckoutUsingPostResponse), nil
	}
}

// ShowUndoCheckoutUsingPostInvoker XDM_撤销检出VersionModel
func (c *IDMEClassicAPIClient) ShowUndoCheckoutUsingPostInvoker(request *model.ShowUndoCheckoutUsingPostRequest) *ShowUndoCheckoutUsingPostInvoker {
	requestDef := GenReqDefForShowUndoCheckoutUsingPost()
	return &ShowUndoCheckoutUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpdateAndCheckinUsingPost XDM_更新并检入VersionModel
//
// 根据传入字段更新版本对象并根据主对象ID检入对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUpdateAndCheckinUsingPost(request *model.ShowUpdateAndCheckinUsingPostRequest) (*model.ShowUpdateAndCheckinUsingPostResponse, error) {
	requestDef := GenReqDefForShowUpdateAndCheckinUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpdateAndCheckinUsingPostResponse), nil
	}
}

// ShowUpdateAndCheckinUsingPostInvoker XDM_更新并检入VersionModel
func (c *IDMEClassicAPIClient) ShowUpdateAndCheckinUsingPostInvoker(request *model.ShowUpdateAndCheckinUsingPostRequest) *ShowUpdateAndCheckinUsingPostInvoker {
	requestDef := GenReqDefForShowUpdateAndCheckinUsingPost()
	return &ShowUpdateAndCheckinUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpdateByAdminUsingPost XDM_管理员更新对象VersionModel指定属性
//
// 以管理员身份更新指定版本实例上的基础信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUpdateByAdminUsingPost(request *model.ShowUpdateByAdminUsingPostRequest) (*model.ShowUpdateByAdminUsingPostResponse, error) {
	requestDef := GenReqDefForShowUpdateByAdminUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpdateByAdminUsingPostResponse), nil
	}
}

// ShowUpdateByAdminUsingPostInvoker XDM_管理员更新对象VersionModel指定属性
func (c *IDMEClassicAPIClient) ShowUpdateByAdminUsingPostInvoker(request *model.ShowUpdateByAdminUsingPostRequest) *ShowUpdateByAdminUsingPostInvoker {
	requestDef := GenReqDefForShowUpdateByAdminUsingPost()
	return &ShowUpdateByAdminUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpdateByConditionUsingPost 根据指定条件更新实例
//
// 根据指定条件更新指定模型的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUpdateByConditionUsingPost(request *model.ShowUpdateByConditionUsingPostRequest) (*model.ShowUpdateByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForShowUpdateByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpdateByConditionUsingPostResponse), nil
	}
}

// ShowUpdateByConditionUsingPostInvoker 根据指定条件更新实例
func (c *IDMEClassicAPIClient) ShowUpdateByConditionUsingPostInvoker(request *model.ShowUpdateByConditionUsingPostRequest) *ShowUpdateByConditionUsingPostInvoker {
	requestDef := GenReqDefForShowUpdateByConditionUsingPost()
	return &ShowUpdateByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpdateUsingPost 更新实例
//
// 更新指定数据模型中的一个实例数据。如果实例的唯一编码不存在，则不做任何更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowUpdateUsingPost(request *model.ShowUpdateUsingPostRequest) (*model.ShowUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForShowUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpdateUsingPostResponse), nil
	}
}

// ShowUpdateUsingPostInvoker 更新实例
func (c *IDMEClassicAPIClient) ShowUpdateUsingPostInvoker(request *model.ShowUpdateUsingPostRequest) *ShowUpdateUsingPostInvoker {
	requestDef := GenReqDefForShowUpdateUsingPost()
	return &ShowUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
