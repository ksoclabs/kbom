package cmd

import (
	"bytes"
	"testing"

	"github.com/ksoclabs/kbom/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	mock := &stdoutMock{buf: bytes.Buffer{}}
	out = mock

	config.AppName = "kbom"
	config.AppVersion = "1.0.0"
	config.BuildTime = "2021-01-01T00:00:00Z"
	config.LastCommitHash = "1234567890"

	err := runPrintVersion(nil, []string{})
	assert.NoError(t, err)

	assert.Equal(t, expectedVersion, mock.buf.String())
}

var expectedVersion = `kbom version 1.0.0
build date: 2021-01-01T00:00:00Z
commit: 1234567890

https://github.com/ksoclabs/kbom
`
