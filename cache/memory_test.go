package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemory(t *testing.T) {
	as := assert.New(t)

	m := NewMemory()

	{
		as.Nil(m.Get("a"))
		as.False(m.IsExist("a"))
		as.Nil(m.Delete(""))
	}

	{
		as.Nil(m.Set("a", "hhh", time.Second))
		as.Equal("hhh", m.Get("a"))
		as.True(m.IsExist("a"))
		as.Nil(m.Delete("a"))
		as.False(m.IsExist("a"))
	}

	{
		as.Nil(m.Set("a", "hhh", time.Second))

		time.Sleep(time.Second)

		as.Nil(m.Get("a"))
		as.False(m.IsExist("a"))
		as.Nil(m.Delete("a"))

	}

	{
		for i := 0; i < 100; i++ {
			as.Nil(m.Set("a", "hhh", time.Second))
			as.Nil(m.Delete("a"))
		}
	}
}
