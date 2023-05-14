package projector_test

import (
	"testing"

	projector_config "github.com/zolwiastyl/submarine/pkg/config"
	"github.com/zolwiastyl/submarine/pkg/projector"
)

func handleGetValueTest(t *testing.T, projector *projector.Projector, key string, expected string) {
	value, ok := projector.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", expected)
	}
	handleDifferentValueTest(t, value, expected)
}

func TestGetValue(t *testing.T) {
	data := GetData()
	projector := GetProjector("/foo/bar", data)
	handleGetValueTest(t, projector, "key2", "topValue")
	handleGetValueTest(t, projector, "key1", "value3")
	handleGetValueTest(t, projector, "key3", "isGreat")
}

// It will throw error if value is different than desired
func handleDifferentValueTest(t *testing.T, value string, expected string) {
	if value != expected {
		t.Errorf("expected to find %v but received %v", expected, value)
	}
}

func TestSetValue(t *testing.T) {
	data := GetData()
	projector := GetProjector("/foo/bar", data)
	handleGetValueTest(t, projector, "key2", "topValue")
	projector.SetValue("key2", "newValue")
	handleGetValueTest(t, projector, "key2", "newValue")

	handleGetValueTest(t, projector, "key3", "isGreat")
	projector.SetValue("key3", "newValue")
	handleGetValueTest(t, projector, "key3", "newValue")
}

func TestGetValueAll(t *testing.T) {
	data := GetData()
	projector := GetProjector("/foo/bar", data)
	out := projector.GetValueAll()
	if len(out) != 3 {
		t.Errorf("expected to find 3 values but found %v", len(out))
	}
	handleDifferentValueTest(t, out["key1"], "value3")
	handleDifferentValueTest(t, out["key2"], "topValue")
	handleDifferentValueTest(t, out["key3"], "isGreat")

	projector = GetProjector("/foo", data)
	out = projector.GetValueAll()
	if len(out) != 2 {
		t.Errorf("expected to find 2 values but found %v", len(out))
	}
	handleDifferentValueTest(t, out["key1"], "value2")
	handleDifferentValueTest(t, out["key3"], "isGreat")

	projector = GetProjector("/", data)
	out = projector.GetValueAll()
	if len(out) != 2 {
		t.Errorf("expected to find 2 values but found %v", len(out))
	}
	handleDifferentValueTest(t, out["key1"], "value1")
	handleDifferentValueTest(t, out["key3"], "isGreat")
}

func TestRemoveValue(t *testing.T) {
	data := GetData()
	projector := GetProjector("/foo/bar", data)
	key := "key2"

	handleGetValueTest(t, projector, key, "topValue")

	projector.RemoveValue(key)
	_, ok := projector.GetValue(key)
	if ok {
		t.Errorf("expected to not find value \"%v\"", key)
	}

	key = "key1"
	handleGetValueTest(t, projector, key, "value3")
	projector.RemoveValue(key)
	handleGetValueTest(t, projector, key, "value2")
}

func GetData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"key1": "value1",
				"key3": "isGreat",
			},
			"/foo": {
				"key1": "value2",
			},
			"/foo/bar": {
				"key1": "value3",
				"key2": "topValue",
			},
		},
	}
}

func GetProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector_config.Config{
			Args:       []string{},
			Operation:  projector_config.Print,
			Pwd:        pwd,
			ConfigPath: "",
		},
		data,
	)
}
