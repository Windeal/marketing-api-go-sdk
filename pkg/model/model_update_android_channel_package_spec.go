/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 更新应用宝渠道包接口任务所需条件
type UpdateAndroidChannelPackageSpec struct {
	ChannelPackageId string `json:"channel_package_id,omitempty"`
	DownloadUrl      string `json:"download_url,omitempty"`
}