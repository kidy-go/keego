// config_test.go kee > 2021/02/01

package conf

import (
	"github.com/spf13/cast"
	"keego/test"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := LoadEnv(".")
	test.Equal(t, cfg.Get("APP_ENV", "no val"), "developer")
	test.Equal(t, cfg.Get("app_env", "no val"), "developer")
	test.UnEqual(t, cfg.Get("app.env", "no val"), "developer")

	test.Equal(t, cast.ToBool(cfg.Get("APP_DEBUG")), true)

	cfg.Set("app.env", "dev")
	test.Equal(t, cfg.Get("app.env"), "dev")

	ycfg := LoadYaml("./", "test.yaml")
	test.Equal(t, ycfg.Get("version"), "3.3")
	test.Equal(t, ycfg.Get("services.test.image"), "test:latest")
	test.Equal(t, ycfg.Get("services.test.name", "dev"), "dev")
	ycfg.Set("services.test.name", "test")
	test.UnEqual(t, ycfg.Get("services.test.name", "dev"), "dev")

	ycfg.SaveAs("./test2.yaml", true)
}
