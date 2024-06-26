package pikpak_share

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	driver.RootID
	Username     string `json:"username" required:"true"`
	Password     string `json:"password" required:"true"`
	ShareId      string `json:"share_id" required:"true"`
	SharePwd     string `json:"share_pwd"`
	ClientID     string `json:"client_id" required:"true" default:"YNxT9w7GMdWvEOKa"`
	ClientSecret string `json:"client_secret" required:"true" default:"dbw2OtmVEeuUvIptb1Coyg"`
}

var config = driver.Config{
	Name:        "PikPakShare",
	LocalSort:   true,
	NoUpload:    true,
	DefaultRoot: "",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &PikPakShare{}
	})
}
