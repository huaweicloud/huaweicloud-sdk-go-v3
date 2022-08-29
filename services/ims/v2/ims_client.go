package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
)

type ImsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewImsClient(hcClient *http_client.HcHttpClient) *ImsClient {
	return &ImsClient{HcClient: hcClient}
}

func ImsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddImageTag 添加镜像标签
//
// 该接口用于为指定镜像添加或更新指定的单个标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) AddImageTag(request *model.AddImageTagRequest) (*model.AddImageTagResponse, error) {
	requestDef := GenReqDefForAddImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddImageTagResponse), nil
	}
}

// AddImageTagInvoker 添加镜像标签
func (c *ImsClient) AddImageTagInvoker(request *model.AddImageTagRequest) *AddImageTagInvoker {
	requestDef := GenReqDefForAddImageTag()
	return &AddImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddMembers 批量添加镜像成员
//
// 该接口为扩展接口，主要用于镜像共享时用户将多个镜像共享给多个用户。
// 该接口为异步接口，返回job_id说明任务下发成功，查询异步任务状态，如果是success说明任务执行成功，如果是failed说明任务执行失败。如何查询异步任务，请参见异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) BatchAddMembers(request *model.BatchAddMembersRequest) (*model.BatchAddMembersResponse, error) {
	requestDef := GenReqDefForBatchAddMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddMembersResponse), nil
	}
}

// BatchAddMembersInvoker 批量添加镜像成员
func (c *ImsClient) BatchAddMembersInvoker(request *model.BatchAddMembersRequest) *BatchAddMembersInvoker {
	requestDef := GenReqDefForBatchAddMembers()
	return &BatchAddMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddOrDeleteTags 批量添加删除镜像标签
//
// 该接口用于为指定镜像批量添加/更新、删除标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) BatchAddOrDeleteTags(request *model.BatchAddOrDeleteTagsRequest) (*model.BatchAddOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchAddOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddOrDeleteTagsResponse), nil
	}
}

// BatchAddOrDeleteTagsInvoker 批量添加删除镜像标签
func (c *ImsClient) BatchAddOrDeleteTagsInvoker(request *model.BatchAddOrDeleteTagsRequest) *BatchAddOrDeleteTagsInvoker {
	requestDef := GenReqDefForBatchAddOrDeleteTags()
	return &BatchAddOrDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMembers 批量删除镜像成员
//
// 该接口为扩展接口，主要用于取消镜像共享。
// 该接口为异步接口，返回job_id说明任务下发成功，查询异步任务状态，如果是success说明任务执行成功，如果是failed说明任务执行失败。如何查询异步任务，请参见异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) BatchDeleteMembers(request *model.BatchDeleteMembersRequest) (*model.BatchDeleteMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersResponse), nil
	}
}

// BatchDeleteMembersInvoker 批量删除镜像成员
func (c *ImsClient) BatchDeleteMembersInvoker(request *model.BatchDeleteMembersRequest) *BatchDeleteMembersInvoker {
	requestDef := GenReqDefForBatchDeleteMembers()
	return &BatchDeleteMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateMembers 批量更新镜像成员状态
//
// 该接口为扩展接口，主要用于用户接受或者拒绝多个共享镜像时批量更新镜像成员的状态。
// 该接口为异步接口，返回job_id说明任务下发成功，查询异步任务状态，如果是success说明任务执行成功，如果是failed说明任务执行失败。如何查询异步任务，请参见异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) BatchUpdateMembers(request *model.BatchUpdateMembersRequest) (*model.BatchUpdateMembersResponse, error) {
	requestDef := GenReqDefForBatchUpdateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateMembersResponse), nil
	}
}

// BatchUpdateMembersInvoker 批量更新镜像成员状态
func (c *ImsClient) BatchUpdateMembersInvoker(request *model.BatchUpdateMembersRequest) *BatchUpdateMembersInvoker {
	requestDef := GenReqDefForBatchUpdateMembers()
	return &BatchUpdateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyImageCrossRegion 跨Region复制镜像
//
// 该接口为扩展接口，用户在一个区域制作的私有镜像，可以通过跨Region复制镜像将镜像复制到其他区域，在其他区域发放相同类型的云服务器，帮助用户实现区域间的业务迁移。
// 该接口为异步接口，返回job_id说明任务下发成功，查询异步任务状态，如果是success说明任务执行成功，如果是failed说明任务执行失败。
// 如何查询异步任务，请参见异步任务进度查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CopyImageCrossRegion(request *model.CopyImageCrossRegionRequest) (*model.CopyImageCrossRegionResponse, error) {
	requestDef := GenReqDefForCopyImageCrossRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyImageCrossRegionResponse), nil
	}
}

// CopyImageCrossRegionInvoker 跨Region复制镜像
func (c *ImsClient) CopyImageCrossRegionInvoker(request *model.CopyImageCrossRegionRequest) *CopyImageCrossRegionInvoker {
	requestDef := GenReqDefForCopyImageCrossRegion()
	return &CopyImageCrossRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyImageInRegion Region内复制镜像
//
// 该接口为扩展接口，主要用于用户将一个已有镜像复制为另一个镜像。复制镜像时，可以更改镜像的加密等属性，以满足不同的场景。
// 该接口为异步接口，返回job_id说明任务下发成功，查询异步任务状态，如果是success说明任务执行成功，如果是failed说明任务执行失败。如何查询异步任务，请参见异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CopyImageInRegion(request *model.CopyImageInRegionRequest) (*model.CopyImageInRegionResponse, error) {
	requestDef := GenReqDefForCopyImageInRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyImageInRegionResponse), nil
	}
}

// CopyImageInRegionInvoker Region内复制镜像
func (c *ImsClient) CopyImageInRegionInvoker(request *model.CopyImageInRegionRequest) *CopyImageInRegionInvoker {
	requestDef := GenReqDefForCopyImageInRegion()
	return &CopyImageInRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataImage 使用外部镜像文件制作数据镜像
//
// 使用上传至OBS桶中的外部数据卷镜像文件制作数据镜像。作为异步接口，调用成功，只是说明后台收到了制作请求，镜像是否制作成功需要通过异步任务查询接口查询该任务的执行状态。具体请参考异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CreateDataImage(request *model.CreateDataImageRequest) (*model.CreateDataImageResponse, error) {
	requestDef := GenReqDefForCreateDataImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataImageResponse), nil
	}
}

// CreateDataImageInvoker 使用外部镜像文件制作数据镜像
func (c *ImsClient) CreateDataImageInvoker(request *model.CreateDataImageRequest) *CreateDataImageInvoker {
	requestDef := GenReqDefForCreateDataImage()
	return &CreateDataImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImage 制作镜像
//
// 本接口用于制作私有镜像，支持：
// - 使用云服务器制作私有镜像。
// - 使用上传至OBS桶中的外部镜像文件制作私有镜像。
// - 使用数据卷制作系统盘镜像。
//
// 作为异步接口，调用成功，只是说明云平台收到了制作请求，镜像是否制作成功需要通过异步任务查询接口查询该任务的执行状态，具体请参考异步任务查询。
//
// 不同场景必选参数说明：
//
// - 使用云服务器制作镜像时的请求的必选参数：name,instance_id。
// - 使用上传至OBS桶中的外部镜像文件时的请求必选参数：name,image_url,min_disk。
// - 使用数据卷制作系统盘镜像时的请求必选参数：name,volume_id,os_version
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

// CreateImageInvoker 制作镜像
func (c *ImsClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateTags 增加或修改标签
//
// 该接口主要用于为某个镜像增加或修改一个自定义标签。通过自定义标签，用户可以将镜像进行分类。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CreateOrUpdateTags(request *model.CreateOrUpdateTagsRequest) (*model.CreateOrUpdateTagsResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateTagsResponse), nil
	}
}

// CreateOrUpdateTagsInvoker 增加或修改标签
func (c *ImsClient) CreateOrUpdateTagsInvoker(request *model.CreateOrUpdateTagsRequest) *CreateOrUpdateTagsInvoker {
	requestDef := GenReqDefForCreateOrUpdateTags()
	return &CreateOrUpdateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWholeImage 制作整机镜像
//
// 使用云服务器或者云服务器备份制作整机镜像。作为异步接口，调用成功，只是说明后台收到了制作整机镜像的请求，镜像是否制作成功需要通过异步任务查询接口查询该任务的执行状态，具体请参考异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) CreateWholeImage(request *model.CreateWholeImageRequest) (*model.CreateWholeImageResponse, error) {
	requestDef := GenReqDefForCreateWholeImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWholeImageResponse), nil
	}
}

// CreateWholeImageInvoker 制作整机镜像
func (c *ImsClient) CreateWholeImageInvoker(request *model.CreateWholeImageRequest) *CreateWholeImageInvoker {
	requestDef := GenReqDefForCreateWholeImage()
	return &CreateWholeImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImageTag 删除镜像标签
//
// 该接口用于为镜像删除指定的标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) DeleteImageTag(request *model.DeleteImageTagRequest) (*model.DeleteImageTagResponse, error) {
	requestDef := GenReqDefForDeleteImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageTagResponse), nil
	}
}

// DeleteImageTagInvoker 删除镜像标签
func (c *ImsClient) DeleteImageTagInvoker(request *model.DeleteImageTagRequest) *DeleteImageTagInvoker {
	requestDef := GenReqDefForDeleteImageTag()
	return &DeleteImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportImage 导出镜像
//
// 该接口为扩展接口，用于用户将自己的私有镜像导出到指定的OBS桶中。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ExportImage(request *model.ExportImageRequest) (*model.ExportImageResponse, error) {
	requestDef := GenReqDefForExportImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportImageResponse), nil
	}
}

// ExportImageInvoker 导出镜像
func (c *ImsClient) ExportImageInvoker(request *model.ExportImageRequest) *ExportImageInvoker {
	requestDef := GenReqDefForExportImage()
	return &ExportImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportImageQuick 镜像文件快速导入
//
// 使用上传至OBS桶中的超大外部镜像文件制作私有镜像，目前仅支持RAW或ZVHD2格式镜像文件。且要求镜像文件大小不能超过1TB。
// 由于快速导入功能要求提前转换镜像文件格式为RAW或ZVHD2格式，因此镜像文件小于128GB时推荐您优先使用常规的创建私有镜像的方式。
// 作为异步接口，调用成功，只是说明后台收到了制作请求，镜像是否制作成功需要通过异步任务查询接口查询该任务的执行状态，具体请参考异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ImportImageQuick(request *model.ImportImageQuickRequest) (*model.ImportImageQuickResponse, error) {
	requestDef := GenReqDefForImportImageQuick()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportImageQuickResponse), nil
	}
}

// ImportImageQuickInvoker 镜像文件快速导入
func (c *ImsClient) ImportImageQuickInvoker(request *model.ImportImageQuickRequest) *ImportImageQuickInvoker {
	requestDef := GenReqDefForImportImageQuick()
	return &ImportImageQuickInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageByTags 按标签查询镜像
//
// 该接口用于按标签或其他条件对镜像进行过滤或者计数使用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListImageByTags(request *model.ListImageByTagsRequest) (*model.ListImageByTagsResponse, error) {
	requestDef := GenReqDefForListImageByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageByTagsResponse), nil
	}
}

// ListImageByTagsInvoker 按标签查询镜像
func (c *ImsClient) ListImageByTagsInvoker(request *model.ListImageByTagsRequest) *ListImageByTagsInvoker {
	requestDef := GenReqDefForListImageByTags()
	return &ListImageByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageTags 查询镜像标签
//
// 该接口用于为查询指定镜像上的所有标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListImageTags(request *model.ListImageTagsRequest) (*model.ListImageTagsResponse, error) {
	requestDef := GenReqDefForListImageTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageTagsResponse), nil
	}
}

// ListImageTagsInvoker 查询镜像标签
func (c *ImsClient) ListImageTagsInvoker(request *model.ListImageTagsRequest) *ListImageTagsInvoker {
	requestDef := GenReqDefForListImageTags()
	return &ListImageTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImages 查询镜像列表
//
// 根据不同条件查询镜像列表信息。
// 可以在URI后面用‘?’和‘&amp;’添加不同的查询条件组合，请参考请求样例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

// ListImagesInvoker 查询镜像列表
func (c *ImsClient) ListImagesInvoker(request *model.ListImagesRequest) *ListImagesInvoker {
	requestDef := GenReqDefForListImages()
	return &ListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImagesTags 查询租户所有镜像标签
//
// 该接口用于为查询租户的所有镜像上的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListImagesTags(request *model.ListImagesTagsRequest) (*model.ListImagesTagsResponse, error) {
	requestDef := GenReqDefForListImagesTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesTagsResponse), nil
	}
}

// ListImagesTagsInvoker 查询租户所有镜像标签
func (c *ImsClient) ListImagesTagsInvoker(request *model.ListImagesTagsRequest) *ListImagesTagsInvoker {
	requestDef := GenReqDefForListImagesTags()
	return &ListImagesTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOsVersions 查询镜像支持的OS列表
//
// 查询当前区域弹性云服务器的OS兼容性列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListOsVersions(request *model.ListOsVersionsRequest) (*model.ListOsVersionsResponse, error) {
	requestDef := GenReqDefForListOsVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOsVersionsResponse), nil
	}
}

// ListOsVersionsInvoker 查询镜像支持的OS列表
func (c *ImsClient) ListOsVersionsInvoker(request *model.ListOsVersionsRequest) *ListOsVersionsInvoker {
	requestDef := GenReqDefForListOsVersions()
	return &ListOsVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 按条件查询租户镜像标签列表
//
// 根据不同条件查询镜像标签列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 按条件查询租户镜像标签列表
func (c *ImsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterImage 注册镜像
//
// 该接口用于将镜像文件注册为云平台未初始化的私有镜像。
// 使用该接口注册镜像的具体步骤如下：
// 将镜像文件上传到OBS个人桶中。具体操作请参见《对象存储服务客户端指南（OBS Browser）》或《对象存储服务API参考》。
// 使用创建镜像元数据接口创建镜像元数据。调用成功后，保存该镜像的ID。创建镜像元数据请参考创建镜像元数据（OpenStack原生）。
// 根据2得到的镜像ID，使用注册镜像接口注册OBS桶中的镜像文件。
// 注册镜像接口作为异步接口，调用成功后，说明后台收到了注册请求。需要根据镜像ID查询该镜像状态验证镜像注册是否成功。当镜像状态变为“active”时，表示镜像注册成功。
// 如何查询异步任务，请参见异步任务查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) RegisterImage(request *model.RegisterImageRequest) (*model.RegisterImageResponse, error) {
	requestDef := GenReqDefForRegisterImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImageResponse), nil
	}
}

// RegisterImageInvoker 注册镜像
func (c *ImsClient) RegisterImageInvoker(request *model.RegisterImageRequest) *RegisterImageInvoker {
	requestDef := GenReqDefForRegisterImage()
	return &RegisterImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageQuota 查询镜像配额
//
// 该接口为扩展接口，主要用于查询租户在当前Region的私有镜像的配额数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ShowImageQuota(request *model.ShowImageQuotaRequest) (*model.ShowImageQuotaResponse, error) {
	requestDef := GenReqDefForShowImageQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageQuotaResponse), nil
	}
}

// ShowImageQuotaInvoker 查询镜像配额
func (c *ImsClient) ShowImageQuotaInvoker(request *model.ShowImageQuotaRequest) *ShowImageQuotaInvoker {
	requestDef := GenReqDefForShowImageQuota()
	return &ShowImageQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询job状态
//
// 该接口为扩展接口，主要用于查询异步接口执行情况，比如查询导出镜像任务的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询job状态
func (c *ImsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobProgress 异步任务进度查询
//
// 该接口为扩展接口，主要用于查询异步任务进度。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ShowJobProgress(request *model.ShowJobProgressRequest) (*model.ShowJobProgressResponse, error) {
	requestDef := GenReqDefForShowJobProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobProgressResponse), nil
	}
}

// ShowJobProgressInvoker 异步任务进度查询
func (c *ImsClient) ShowJobProgressInvoker(request *model.ShowJobProgressRequest) *ShowJobProgressInvoker {
	requestDef := GenReqDefForShowJobProgress()
	return &ShowJobProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImage 更新镜像信息
//
// 更新镜像信息接口，主要用于镜像属性的修改。当前仅支持可用（active）状态的镜像更新相关信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) UpdateImage(request *model.UpdateImageRequest) (*model.UpdateImageResponse, error) {
	requestDef := GenReqDefForUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageResponse), nil
	}
}

// UpdateImageInvoker 更新镜像信息
func (c *ImsClient) UpdateImageInvoker(request *model.UpdateImageRequest) *UpdateImageInvoker {
	requestDef := GenReqDefForUpdateImage()
	return &UpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersions 查询版本列表（OpenStack原生）
//
// 查询API的版本信息列表，包括API的版本兼容性、域名信息等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

// ListVersionsInvoker 查询版本列表（OpenStack原生）
func (c *ImsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 查询版本列表（OpenStack原生）
//
// 查询API的版本信息列表，包括API的版本兼容性、域名信息等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 查询版本列表（OpenStack原生）
func (c *ImsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceAddImageMember 添加镜像成员（OpenStack原生）
//
// 用户共享镜像给其他用户时，使用该接口向该镜像成员中添加接受镜像用户的项目ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceAddImageMember(request *model.GlanceAddImageMemberRequest) (*model.GlanceAddImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceAddImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceAddImageMemberResponse), nil
	}
}

// GlanceAddImageMemberInvoker 添加镜像成员（OpenStack原生）
func (c *ImsClient) GlanceAddImageMemberInvoker(request *model.GlanceAddImageMemberRequest) *GlanceAddImageMemberInvoker {
	requestDef := GenReqDefForGlanceAddImageMember()
	return &GlanceAddImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceCreateImageMetadata 创建镜像元数据（OpenStack原生）
//
// 创建镜像元数据。调用创建镜像元数据接口成功后，只是创建了镜像的元数据，镜像对应的实际镜像文件并不存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceCreateImageMetadata(request *model.GlanceCreateImageMetadataRequest) (*model.GlanceCreateImageMetadataResponse, error) {
	requestDef := GenReqDefForGlanceCreateImageMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceCreateImageMetadataResponse), nil
	}
}

// GlanceCreateImageMetadataInvoker 创建镜像元数据（OpenStack原生）
func (c *ImsClient) GlanceCreateImageMetadataInvoker(request *model.GlanceCreateImageMetadataRequest) *GlanceCreateImageMetadataInvoker {
	requestDef := GenReqDefForGlanceCreateImageMetadata()
	return &GlanceCreateImageMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceCreateTag 增加标签（OpenStack原生）
//
// 该接口主要用于为某个镜像添加一个自定义标签。通过自定义标签，用户可以将镜像进行分类。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceCreateTag(request *model.GlanceCreateTagRequest) (*model.GlanceCreateTagResponse, error) {
	requestDef := GenReqDefForGlanceCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceCreateTagResponse), nil
	}
}

// GlanceCreateTagInvoker 增加标签（OpenStack原生）
func (c *ImsClient) GlanceCreateTagInvoker(request *model.GlanceCreateTagRequest) *GlanceCreateTagInvoker {
	requestDef := GenReqDefForGlanceCreateTag()
	return &GlanceCreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceDeleteImage 删除镜像（OpenStack原生）
//
// 该接口主要用于删除镜像，用户可以通过该接口将自己的私有镜像删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceDeleteImage(request *model.GlanceDeleteImageRequest) (*model.GlanceDeleteImageResponse, error) {
	requestDef := GenReqDefForGlanceDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteImageResponse), nil
	}
}

// GlanceDeleteImageInvoker 删除镜像（OpenStack原生）
func (c *ImsClient) GlanceDeleteImageInvoker(request *model.GlanceDeleteImageRequest) *GlanceDeleteImageInvoker {
	requestDef := GenReqDefForGlanceDeleteImage()
	return &GlanceDeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceDeleteImageMember 删除指定的镜像成员（OpenStack原生）
//
// 该接口用于取消对某个用户的镜像共享。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceDeleteImageMember(request *model.GlanceDeleteImageMemberRequest) (*model.GlanceDeleteImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceDeleteImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteImageMemberResponse), nil
	}
}

// GlanceDeleteImageMemberInvoker 删除指定的镜像成员（OpenStack原生）
func (c *ImsClient) GlanceDeleteImageMemberInvoker(request *model.GlanceDeleteImageMemberRequest) *GlanceDeleteImageMemberInvoker {
	requestDef := GenReqDefForGlanceDeleteImageMember()
	return &GlanceDeleteImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceDeleteTag 删除标签（OpenStack原生）
//
// 该接口主要用于删除某个镜像的自定义标签，通过该接口，用户可以将私有镜像中一些不用的标签删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceDeleteTag(request *model.GlanceDeleteTagRequest) (*model.GlanceDeleteTagResponse, error) {
	requestDef := GenReqDefForGlanceDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteTagResponse), nil
	}
}

// GlanceDeleteTagInvoker 删除标签（OpenStack原生）
func (c *ImsClient) GlanceDeleteTagInvoker(request *model.GlanceDeleteTagRequest) *GlanceDeleteTagInvoker {
	requestDef := GenReqDefForGlanceDeleteTag()
	return &GlanceDeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceListImageMemberSchemas 查询镜像成员列表视图（OpenStack原生）
//
// 该接口主要用于查询镜像成员列表视图，通过视图，用户可以了解到镜像成员包含哪些属性，同时也可以了解每个属性的数据类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceListImageMemberSchemas(request *model.GlanceListImageMemberSchemasRequest) (*model.GlanceListImageMemberSchemasResponse, error) {
	requestDef := GenReqDefForGlanceListImageMemberSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageMemberSchemasResponse), nil
	}
}

// GlanceListImageMemberSchemasInvoker 查询镜像成员列表视图（OpenStack原生）
func (c *ImsClient) GlanceListImageMemberSchemasInvoker(request *model.GlanceListImageMemberSchemasRequest) *GlanceListImageMemberSchemasInvoker {
	requestDef := GenReqDefForGlanceListImageMemberSchemas()
	return &GlanceListImageMemberSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceListImageMembers 获取镜像成员列表（OpenStack原生）
//
// 该接口用于共享镜像过程中，获取接受该镜像的成员列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceListImageMembers(request *model.GlanceListImageMembersRequest) (*model.GlanceListImageMembersResponse, error) {
	requestDef := GenReqDefForGlanceListImageMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageMembersResponse), nil
	}
}

// GlanceListImageMembersInvoker 获取镜像成员列表（OpenStack原生）
func (c *ImsClient) GlanceListImageMembersInvoker(request *model.GlanceListImageMembersRequest) *GlanceListImageMembersInvoker {
	requestDef := GenReqDefForGlanceListImageMembers()
	return &GlanceListImageMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceListImageSchemas 查询镜像列表视图（OpenStack原生）
//
// 该接口主要用于查询镜像列表视图，通过该接口用户可以了解到镜像列表的详细情况和数据结构。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceListImageSchemas(request *model.GlanceListImageSchemasRequest) (*model.GlanceListImageSchemasResponse, error) {
	requestDef := GenReqDefForGlanceListImageSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageSchemasResponse), nil
	}
}

// GlanceListImageSchemasInvoker 查询镜像列表视图（OpenStack原生）
func (c *ImsClient) GlanceListImageSchemasInvoker(request *model.GlanceListImageSchemasRequest) *GlanceListImageSchemasInvoker {
	requestDef := GenReqDefForGlanceListImageSchemas()
	return &GlanceListImageSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceListImages 查询镜像列表（OpenStack原生）
//
// 获取镜像列表。
// 使用本接口查询镜像列表时，需要使用分页查询才能返回全部的镜像列表。
// 分页说明
// 分页是指返回一组镜像的一个子集，在返回的时候会存在下个子集的链接和首个子集的链接，默认返回的子集中数量为25，用户也可以通过使用limit和marker两个参数自己分页，指定返回子集中需要返回的数量。
// 响应中的参数first是查询首页的URL。next是查询下一页的URL。当查询镜像列表最后一页时，不存在next。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceListImages(request *model.GlanceListImagesRequest) (*model.GlanceListImagesResponse, error) {
	requestDef := GenReqDefForGlanceListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImagesResponse), nil
	}
}

// GlanceListImagesInvoker 查询镜像列表（OpenStack原生）
func (c *ImsClient) GlanceListImagesInvoker(request *model.GlanceListImagesRequest) *GlanceListImagesInvoker {
	requestDef := GenReqDefForGlanceListImages()
	return &GlanceListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceShowImage 查询镜像详情（OpenStack原生）
//
// 查询单个镜像详情，用户可以通过该接口查询单个私有或者公共镜像的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceShowImage(request *model.GlanceShowImageRequest) (*model.GlanceShowImageResponse, error) {
	requestDef := GenReqDefForGlanceShowImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageResponse), nil
	}
}

// GlanceShowImageInvoker 查询镜像详情（OpenStack原生）
func (c *ImsClient) GlanceShowImageInvoker(request *model.GlanceShowImageRequest) *GlanceShowImageInvoker {
	requestDef := GenReqDefForGlanceShowImage()
	return &GlanceShowImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceShowImageMember 获取镜像成员详情（OpenStack原生）
//
// 该接口主要用于镜像共享中查询某个镜像成员的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceShowImageMember(request *model.GlanceShowImageMemberRequest) (*model.GlanceShowImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceShowImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageMemberResponse), nil
	}
}

// GlanceShowImageMemberInvoker 获取镜像成员详情（OpenStack原生）
func (c *ImsClient) GlanceShowImageMemberInvoker(request *model.GlanceShowImageMemberRequest) *GlanceShowImageMemberInvoker {
	requestDef := GenReqDefForGlanceShowImageMember()
	return &GlanceShowImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceShowImageMemberSchemas 查询镜像成员视图（OpenStack原生）
//
// 该接口主要用于查询镜像成员视图，通过视图，用户可以了解到镜像成员包含哪些属性，同时也可以了解每个属性的数据类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceShowImageMemberSchemas(request *model.GlanceShowImageMemberSchemasRequest) (*model.GlanceShowImageMemberSchemasResponse, error) {
	requestDef := GenReqDefForGlanceShowImageMemberSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageMemberSchemasResponse), nil
	}
}

// GlanceShowImageMemberSchemasInvoker 查询镜像成员视图（OpenStack原生）
func (c *ImsClient) GlanceShowImageMemberSchemasInvoker(request *model.GlanceShowImageMemberSchemasRequest) *GlanceShowImageMemberSchemasInvoker {
	requestDef := GenReqDefForGlanceShowImageMemberSchemas()
	return &GlanceShowImageMemberSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceShowImageSchemas 查询镜像视图（OpenStack原生）
//
// 该接口主要用于查询镜像视图，通过视图，用户可以了解到镜像包含哪些属性，同时也可以了解每个属性的数据类型等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceShowImageSchemas(request *model.GlanceShowImageSchemasRequest) (*model.GlanceShowImageSchemasResponse, error) {
	requestDef := GenReqDefForGlanceShowImageSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageSchemasResponse), nil
	}
}

// GlanceShowImageSchemasInvoker 查询镜像视图（OpenStack原生）
func (c *ImsClient) GlanceShowImageSchemasInvoker(request *model.GlanceShowImageSchemasRequest) *GlanceShowImageSchemasInvoker {
	requestDef := GenReqDefForGlanceShowImageSchemas()
	return &GlanceShowImageSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceUpdateImage 更新镜像信息（OpenStack原生）
//
// 修改镜像信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceUpdateImage(request *model.GlanceUpdateImageRequest) (*model.GlanceUpdateImageResponse, error) {
	requestDef := GenReqDefForGlanceUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceUpdateImageResponse), nil
	}
}

// GlanceUpdateImageInvoker 更新镜像信息（OpenStack原生）
func (c *ImsClient) GlanceUpdateImageInvoker(request *model.GlanceUpdateImageRequest) *GlanceUpdateImageInvoker {
	requestDef := GenReqDefForGlanceUpdateImage()
	return &GlanceUpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GlanceUpdateImageMember 更新镜像成员状态（OpenStack原生）
//
// 用户接受或者拒绝共享镜像时，使用该接口更新镜像成员的状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ImsClient) GlanceUpdateImageMember(request *model.GlanceUpdateImageMemberRequest) (*model.GlanceUpdateImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceUpdateImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceUpdateImageMemberResponse), nil
	}
}

// GlanceUpdateImageMemberInvoker 更新镜像成员状态（OpenStack原生）
func (c *ImsClient) GlanceUpdateImageMemberInvoker(request *model.GlanceUpdateImageMemberRequest) *GlanceUpdateImageMemberInvoker {
	requestDef := GenReqDefForGlanceUpdateImageMember()
	return &GlanceUpdateImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
