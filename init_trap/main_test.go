package init

import (
	"testing"

	"./api"
	"./config"
	"github.com/stretchr/testify/assert"
)

func TestUseApi(t *testing.T) {
	// init config
	config.Init()

	assert.True(t, api.IsEnabled())
}
