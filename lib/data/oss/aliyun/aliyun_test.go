package aliyun_test

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	//os.Exit(-1)
}

// CreateBucket :
func CreateBucket() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyID := "<yourAccessKeyId>"
	accessKeySecret := "<yourAccessKeySecret>"
	bucketName := "<yourBucketName>"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		handleError(err)
	}
}

// Upload :
func Upload() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyID := "<yourAccessKeyId>"
	accessKeySecret := "<yourAccessKeySecret>"
	bucketName := "<yourBucketName>"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "<yourObjectName>"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "<yourLocalFileName>"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}

// Down :
func Down() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyID := "<yourAccessKeyId>"
	accessKeySecret := "<yourAccessKeySecret>"
	bucketName := "<yourBucketName>"
	// <yourObjectName>从OSS下载文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "<yourObjectName>"
	downloadedFileName := "<yourDownloadedFileName>"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 下载文件。
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		handleError(err)
	}
}

// Delete :
func Delete() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "<yourAccessKeyId>"
	accessKeySecret := "<yourAccessKeySecret>"
	bucketName := "<yourBucketName>"
	// <yourObjectName>表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "<yourObjectName>"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 删除文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		handleError(err)
	}
}
