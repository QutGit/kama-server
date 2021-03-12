package opentracing

import (
	"context"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const (
	accessKey = "4hlbrqpQqXUxD0W65D_bq1UFFivHpoiPmIGB4aQb"
	secretKey = "L4hbhdsSm8dCK_t7e22jVG0qUsKzWnexLrn9FApD"
	bucket    = "zuolinju"
)

func quniuConfig() (cfg storage.Config) {
	cfg = storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false
	return
}

func Upload(file io.Reader, filename string, size int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cig := quniuConfig()
	formUploader := storage.NewFormUploader(&cig)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, file, size, &putExtra)

	if err != nil {
		return "", err
	}

	return ret.Key, nil
}
