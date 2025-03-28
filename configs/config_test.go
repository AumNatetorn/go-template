package configs_test

import (
	"encoding/json"
	"fmt"
	"go-template/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	// given
	configs.Init(".")

	// when
	conf := configs.GetConfig()

	// then
	assert.Equal(t, "go-template", conf.App.Name)

	b, _ := json.Marshal(conf)
	fmt.Println(string(b))
}
