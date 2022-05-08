package upload

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path"
	"snowy-video-serve/global"
	"strings"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.SYS_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.SYS_CONFIG.Qiniu.AccessKey, global.SYS_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.SYS_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()                                                                                                       // 创建文件 defer 关闭
	rand.Seed(time.Now().UnixNano())                                                                                      //随机种子
	fileExt := strings.ToLower(path.Ext(file.Filename))                                                                   // 文件后缀
	fileKey := fmt.Sprintf("%s/%d%d%s", global.SYS_CONFIG.Qiniu.PathPrefix, time.Now().Unix(), rand.Intn(91)+10, fileExt) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.SYS_LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.SYS_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@object: *Qiniu
//@function: UploadLocalFile
//@description: 上传本地文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadLocalFile(filePath string, filename string, directory string) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.SYS_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.SYS_CONFIG.Qiniu.AccessKey, global.SYS_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}
	// 创建文件 defer 关闭
	// fileExt := strings.ToLower(path.Ext(filename))                                                                        // 文件后缀
	fileKey := fmt.Sprintf("%s/%s%s", global.SYS_CONFIG.Qiniu.PathPrefix, directory, filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.PutFile(context.Background(), &ret, upToken, fileKey, filePath, &putExtra)
	if putErr != nil {
		global.SYS_LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.SYS_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.SYS_CONFIG.Qiniu.AccessKey, global.SYS_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.SYS_CONFIG.Qiniu.Bucket, key); err != nil {
		global.SYS_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.SYS_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.SYS_CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.SYS_CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
