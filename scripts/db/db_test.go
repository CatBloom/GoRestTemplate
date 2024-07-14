package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	t.Run("Test DB connection", func(t *testing.T) {
		db := NewDatabase().GetDB()
		defer db.Close()
		assert.NoError(t, db.Ping())
	})
}
