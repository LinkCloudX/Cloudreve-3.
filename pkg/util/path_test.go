package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotPathToStandardPath(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal("/", DotPathToStandardPath(""))
	asserts.Equal("/目录", DotPathToStandardPath("目录"))
	asserts.Equal("/目录/目录2", DotPathToStandardPath("目录,目录2"))
}
