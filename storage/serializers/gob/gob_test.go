package gob

import (
	"testing"
)

func TestGob(t *testing.T) {
	internal.SerializerTester(t, Serializer)
}
