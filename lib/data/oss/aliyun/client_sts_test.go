package aliyun_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
{
  "code": 0,
  "data": {
    "ossConfig": {
      "id": 2,
      "domainId": 1,
      "name": "org-content",
      "description": "机构内容平台",
      "region": "oss-cn-shenzhen",
      "bucket": "dist-res-cn",
      "prefixDir": "content/org/",
      "ossHost": "dist-res-cn.oss-cn-shenzhen.aliyuncs.com",
      "domainHost": "res-cn.makeblock.com"
    },
    "accessKeyId": "STS.NT1hp5tZ5qpbEDg2Tw8nDREnd",
    "accessKeySecret": "FmUnY7fCHTjpK7mL9wFeG1e77w31xPc18bx3zA9tWK1z",
    "stsToken": "CAISlQJ1q6Ft5B2yfSjIr5eEI8qBmYUUxrKJR2LW1lQiNOFovYDFhjz2IH9Lf3BhBOsZtf43mGlZ5/4elqJ0UZAATkvCccZ28sztbZkG5c+T1fau5Jko1be/ewHKeRuZsebWZ+LmNqS/Ht6md1HDkAJq3LL+bk/Mdle5MJqP+/EFA9MMRVv6FxEkYu1bPQx/ssQXGGLMPPK2SH7Qj3HXEVBjt3gz6y524r/txdaHuFiMzg/2y/QJ4p78Od2vaY5nO5FjXtGpm+tsee2D8lYJt0YQq/kv1v0coWqW5YHNOTQLvUXaadiz28Z0MQp0apI9H6N5t/XmnZV6wLeMyN+umksWYroPC3iGFdH4nNGvH/iyJkLNymZpmNCnGoABImvw4LwZl0BCTPnYC6KbHKVWJcxNguyOu7lVH1hcLg0nC1StvAe6LaJkNmrjO8VKzs4rNEc0y6f7ydQswloPo/uUAPRcx8p3Kziml9K7iihSu0FXtAWAYzVjGQKzkh3m42qm6eK2Kdvusewup0BOD0dfGE34LRwi7Qq/bTLJldg=",
    "host": "res-cn.makeblock.com"
  },
  "message": ""
}
*/
func TestSTS_UploadFile(t *testing.T) {
	// comm
	bucketName := "dist-res-cn" //"<yourBucketName>"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "content/org/test.jpeg" // "<yourObjectName>"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "/Users/liampro/Downloads/study-rs/hezi.jpeg" //"<yourLocalFileName>"
	// STS
	yourEndpoint := "http://oss-cn-shenzhen.aliyuncs.com"
	yourAccessKeyId := "STS.NT1hp5tZ5qpbEDg2Tw8nDREnd"
	yourAccessKeySecret := "FmUnY7fCHTjpK7mL9wFeG1e77w31xPc18bx3zA9tWK1z"
	yourSecurityToken := "CAISlQJ1q6Ft5B2yfSjIr5eEI8qBmYUUxrKJR2LW1lQiNOFovYDFhjz2IH9Lf3BhBOsZtf43mGlZ5/4elqJ0UZAATkvCccZ28sztbZkG5c+T1fau5Jko1be/ewHKeRuZsebWZ+LmNqS/Ht6md1HDkAJq3LL+bk/Mdle5MJqP+/EFA9MMRVv6FxEkYu1bPQx/ssQXGGLMPPK2SH7Qj3HXEVBjt3gz6y524r/txdaHuFiMzg/2y/QJ4p78Od2vaY5nO5FjXtGpm+tsee2D8lYJt0YQq/kv1v0coWqW5YHNOTQLvUXaadiz28Z0MQp0apI9H6N5t/XmnZV6wLeMyN+umksWYroPC3iGFdH4nNGvH/iyJkLNymZpmNCnGoABImvw4LwZl0BCTPnYC6KbHKVWJcxNguyOu7lVH1hcLg0nC1StvAe6LaJkNmrjO8VKzs4rNEc0y6f7ydQswloPo/uUAPRcx8p3Kziml9K7iihSu0FXtAWAYzVjGQKzkh3m42qm6eK2Kdvusewup0BOD0dfGE34LRwi7Qq/bTLJldg="
	// 用户拿到STS临时凭证后，通过其中的安全令牌（SecurityToken）和临时访问密钥（AccessKeyId和AccessKeySecret）生成OSSClient。
	// 创建OSSClient实例。
	client, err := oss.New(yourEndpoint, yourAccessKeyId, yourAccessKeySecret, oss.SecurityToken(yourSecurityToken))
	if err != nil {
		fmt.Println("1-Error:", err)
		os.Exit(-1)
		// OSS操作。
	}
	// 获取存储空间。
	bucket, err2 := client.Bucket(bucketName)
	if err2 != nil {
		fmt.Println("2-Error:", err2)
		os.Exit(-1)
	}
	// 上传文件。
	err3 := bucket.PutObjectFromFile(objectName, localFileName)
	if err3 != nil {
		fmt.Println("3-Error:", err3)
		os.Exit(-1)
	}
}
