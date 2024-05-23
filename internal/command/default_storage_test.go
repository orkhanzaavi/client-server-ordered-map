package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultStorage_Set(t *testing.T) {
	t.Run(
		"should set a new item", func(t *testing.T) {
			storage := newDefaultStorage()
			key1 := "key1"
			value1 := "value1"

			storage.Set(key1, value1)
			storage.Set("key2", "value2")

			items := storage.GetAll()

			t.Log("When setting 2 new items")
			t.Log("Then the order should be kept")
			assert.Equal(t, key1, items[0].Key)
			assert.Equal(t, value1, items[0].Value)
			assert.Equal(t, "key2", items[1].Key)
			assert.Equal(t, "value2", items[1].Value)
			t.Log("And count of items should be 2")
			assert.Len(t, items, 2)
		},
	)

	t.Run(
		"should set after deletion the first item", func(t *testing.T) {
			storage := newDefaultStorage()
			key1 := "key1"
			value1 := "value1"

			storage.Set(key1, value1)
			storage.Set("key2", "value2")

			storage.Delete(key1)
			storage.Set("key3", "value3")

			items := storage.GetAll()

			t.Log("When adding a key after deleting the first item of the list")
			t.Log("The new element should be added to the end of the list")
			assert.Equal(t, "key2", items[0].Key)
			assert.Equal(t, "value2", items[0].Value)
			assert.Equal(t, "key3", items[1].Key)
			assert.Equal(t, "value3", items[1].Value)
			t.Log("And count of items should be correct")
			assert.Len(t, items, 2)
		},
	)

	t.Run(
		"should set an item after deletion the last item", func(t *testing.T) {
			storage := newDefaultStorage()

			storage.Set("key1", "value1")
			storage.Set("key2", "value2")

			storage.Delete("key2")
			storage.Set("key3", "value3")

			items := storage.GetAll()

			t.Log("When adding a key after deleting the last item of the list")
			t.Log("The new element should be added to the end of the list")
			assert.Equal(t, "key1", items[0].Key)
			assert.Equal(t, "value1", items[0].Value)
			assert.Equal(t, "key3", items[1].Key)
			assert.Equal(t, "value3", items[1].Value)
			t.Log("And count of items should be correct")
			assert.Len(t, items, 2)
		},
	)

	t.Run(
		"should set an item after deletion the middle item", func(t *testing.T) {
			storage := newDefaultStorage()

			storage.Set("key1", "value1")
			storage.Set("key2", "value2")
			storage.Set("key3", "value3")

			storage.Delete("key2")
			storage.Set("key4", "value4")

			items := storage.GetAll()

			t.Log("When adding a key after deleting the middle item of the list")
			t.Log("The middle item should be removed and the new element should be added to the end of the list")
			assert.Equal(t, "key1", items[0].Key)
			assert.Equal(t, "value1", items[0].Value)
			assert.Equal(t, "key3", items[1].Key)
			assert.Equal(t, "value3", items[1].Value)
			assert.Equal(t, "key4", items[2].Key)
			assert.Equal(t, "value4", items[2].Value)
			t.Log("And count of items should be correct")
			assert.Len(t, items, 3)
		},
	)

	t.Run(
		"should update the value of the existent item", func(t *testing.T) {
			storage := newDefaultStorage()

			storage.Set("key1", "value1")
			storage.Set("key1", "value2")

			items := storage.GetAll()

			t.Log("Given a storage with one item")
			t.Log("When adding a value with the same key")
			t.Log("Then there is only one key with updated value is present")
			assert.Equal(t, "key1", items[0].Key)
			assert.Equal(t, "value2", items[0].Value)
			t.Log("And count of items should be 1")
			assert.Len(t, items, 1)
		},
	)
}
