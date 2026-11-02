package lucide_test

import (
	"io/fs"
	"testing"

	"github.com/go-app-blazar/lucide-font"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbedded(t *testing.T) {
	files, err := fs.Glob(lucide.FS, "*.*")
	require.NoError(t, err)
	assert.NotEmpty(t, files)
	assert.Contains(t, files, "lu-icons.css")
	assert.Contains(t, files, "lu-icons.woff2")
}
