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

/*
{
  "code": 0,
  "data": {
    "ossConfig": {
      "id": 3,
      "domainId": 1,
      "name": "hezi-cms",
      "description": "盒子CMS",
      "region": "oss-cn-shenzhen",
      "bucket": "mblock-how-tos",
      "prefixDir": "hezi/cms/",
      "ossHost": "mblock-how-tos.oss-cn-shenzhen.aliyuncs.com",
      "domainHost": "howtos.makeblock.com"
    },
    "accessKeyId": "STS.NTYyaEZER58iH6Z5AWWKSdJbk",
    "accessKeySecret": "6uGRuRJZd9Cxn1Z129KJBN7YWGfdn81mAqCf8nmsCiWx",
    "stsToken": "CAISkgJ1q6Ft5B2yfSjIr5fsMtvxt5pzgvqCShDr0UECW8R/i4/JiTz2IH9Lf3BhBOsZtf43mGlZ5/4elqFzRo1EAEfBdpPMYE/PaUbzDbDasumZsJYe6vT8a23xZjf/2MjNGZKbKPrWZvaqbX3diyZ32sGUXD6+XlujQ/Lr5Jl8dYZVJCLaCwBLH9BLPABvhdYHPH/KT5aXPwXtn3DbATgn2Ed1gngt7r+kkI/OqEjbhln9wexQp4PuZ4SjdI5sMYV+FNywx6s0VNKYjHAKtEYTrvYu1PEVomeXhLzHXQkNuSfhGvHP79hiIDV+YqUHAKNepJD+76Yn5beKxtmslE0XZb0MAnWHGJrLx9DfCHddiQar1hKRGoABrnSV4mvZQvTWdZUSckIFIeUi+ztG6UP2QjZXkxsVKl6dunzM/j43UsFWoQLvjBet4w1rDREiQdLJcIMCKCTnZMB+5ANwAALxZT8Yay29pgAGtAggpVmrnfwuIRJEyU0gaDYvE5oOxmLJTGS1S2YI+q/K2ijHnYgqPDCM+5iiAVA=",
    "host": "howtos.makeblock.com"
  },
  "message": ""
}
*/
func TestSTS_UploadFile_HeZiCMS(t *testing.T) {
	// comm
	bucketName := "mblock-how-tos" //"<yourBucketName>"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "hezi/cms/test.jpeg" // "<yourObjectName>"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "/Users/liampro/Downloads/study-rs/hezi.jpeg" //"<yourLocalFileName>"
	// STS
	yourEndpoint := "http://oss-cn-shenzhen.aliyuncs.com"
	yourAccessKeyId := "STS.NTYyaEZER58iH6Z5AWWKSdJbk"
	yourAccessKeySecret := "6uGRuRJZd9Cxn1Z129KJBN7YWGfdn81mAqCf8nmsCiWx"
	yourSecurityToken := "CAISkgJ1q6Ft5B2yfSjIr5fsMtvxt5pzgvqCShDr0UECW8R/i4/JiTz2IH9Lf3BhBOsZtf43mGlZ5/4elqFzRo1EAEfBdpPMYE/PaUbzDbDasumZsJYe6vT8a23xZjf/2MjNGZKbKPrWZvaqbX3diyZ32sGUXD6+XlujQ/Lr5Jl8dYZVJCLaCwBLH9BLPABvhdYHPH/KT5aXPwXtn3DbATgn2Ed1gngt7r+kkI/OqEjbhln9wexQp4PuZ4SjdI5sMYV+FNywx6s0VNKYjHAKtEYTrvYu1PEVomeXhLzHXQkNuSfhGvHP79hiIDV+YqUHAKNepJD+76Yn5beKxtmslE0XZb0MAnWHGJrLx9DfCHddiQar1hKRGoABrnSV4mvZQvTWdZUSckIFIeUi+ztG6UP2QjZXkxsVKl6dunzM/j43UsFWoQLvjBet4w1rDREiQdLJcIMCKCTnZMB+5ANwAALxZT8Yay29pgAGtAggpVmrnfwuIRJEyU0gaDYvE5oOxmLJTGS1S2YI+q/K2ijHnYgqPDCM+5iiAVA="
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

func TestSTS_UploadFile_HeZiClient(t *testing.T) {
	// comm
	bucketName := "mblock-how-tos" //"<yourBucketName>"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "hezi/client/test.jpeg" // "<yourObjectName>"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "/Users/liampro/Downloads/study-rs/hezi.jpeg" //"<yourLocalFileName>"
	// STS
	yourEndpoint := "http://oss-cn-shenzhen.aliyuncs.com"
	yourAccessKeyId := "STS.NUz6BMLoL9PrRRAZMec3BzbPq"
	yourAccessKeySecret := "9FXuK87hXL9bSttpk57jMcFrhNP5GirxwjrrKp9cbHVQ"
	yourSecurityToken := "CAISmAJ1q6Ft5B2yfSjIr5bPffj5obBtjpKZUHTwvk0wb7xulaf7kzz2IH9Lf3BhBOsZtf43mGlZ5/4elqJzRo1EAEfAbMZ28szLP+k4+8+T1fau5Jko1be4ewHKeQaZsebWZ+LmNqS/Ht6md1HDkAJq3LL+bk/Mdle5MJqP+/EFA9MMRVv6FxQkYu1bPQx/ssQXGGLMPPK2SH7Qj3HXEVBjt3g+6yN24r/txdaHuFiMzg//wOsSrIe0Ip+7KtVrJ9B/XsW0m+dyd6fHzGkSyWATqPks0/Ido2af5ozMWAkB2XjcbbqIqO8IBRRie603F5RDqPXBjvBisoTR7d+olE0UbLwODH+FHNz5kZqcRPnCMc0iZSIPhKRfcXTlGoABLxU6TsAJYcgTWcUWatIYIVXgHLSniSKtZGzTBf14WFrWVPIuoTA2OBREMLFb8LlP+Hl8eVCWZXwfBVxcnXLxO0rDqOHOI/Ayd/dWRlsne4Nn+Ii8Y12j35uOVnuj2RSW9qe4aGQjw3NHzU+mRxx+oDe0a8iwS/34siJBRlPicpU="
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
