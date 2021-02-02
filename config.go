// Config.go kee > 2021/02/01

package keego

import (
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New(cPath string, cName string, cType string) *Config {
	v := viper.New()
	c := &Config{v}
	c.v.AddConfigPath(cPath)
	c.v.SetConfigName(cName)
	c.v.SetConfigType(cType)

	if e := c.v.ReadInConfig(); e != nil {
		panic(e)
	}

	return c
}

func LoadEnv(envPath string, envName ...string) *Config {
	name := ".env"
	if len(envName) > 0 {
		name = envName[0]
	}
	c := New(envPath, name, "env")
	c.Viper().AutomaticEnv()
	return c
}

func LoadYaml(path, name string) *Config {
	return New(path, name, "yaml")
}

func LoadJson(path, name string) *Config {
	return New(path, name, "json")
}

func LoadToml(path, name string) *Config {
	return New(path, name, "toml")
}

func LoadHCL(path, name string) *Config {
	return New(path, name, "hcl")
}

func (c *Config) Set(key string, value interface{}) {
	c.v.Set(key, value)
}

func (c *Config) Has(key string) bool {
	return c.v.IsSet(key)
}

func (c *Config) Get(key string, defaultValue ...interface{}) interface{} {
	if !c.Has(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return c.v.Get(key)
}

func (c *Config) AddConfigPath(path string) {
	c.v.AddConfigPath(path)
}

func (c *Config) Save(safeWrite ...bool) {
	if len(safeWrite) > 0 {
		if true == safeWrite[0] {
			c.v.SafeWriteConfig()
			return
		}
	}
	c.v.WriteConfig()
}

func (c *Config) SaveAs(path string, safeWrite ...bool) {
	if len(safeWrite) > 0 {
		if true == safeWrite[0] {
			c.v.SafeWriteConfigAs(path)
			return
		}
	}
	c.v.WriteConfigAs(path)
}

func (c *Config) Viper() *viper.Viper {
	return c.v
}
