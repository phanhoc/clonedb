package db

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now().Format("20060102150405"))
}
