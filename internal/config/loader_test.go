package config

import (
	"github.com/Kamran151199/dbmigrate/internal/config/target"
	"github.com/stretchr/testify/require"
	"reflect"
	"strings"
	"testing"
)

func Checker[T target.Target](t *testing.T, conf *T, expected interface{}, err error) {
	confType := reflect.TypeOf(conf)
	t.Logf("Config: %v", *conf)
	t.Logf("Expectancy: %v", expected)
	require.NoError(t, err)

	switch confType.String() {
	case "*target.Mongo":
		if v, ok := expected.(target.Mongo); ok {
			require.NotEmpty(t, conf)
			require.Equal(t, v, *conf)
		} else {
			t.Errorf("Could not conver to type target.Mongo: %v\n", expected)
		}
	case "*target.Postgresql":
		if v, ok := expected.(target.Postgresql); ok {
			require.NotEmpty(t, conf)
			require.Equal(t, v, *conf)
		} else {
			t.Errorf("Could not conver to type target.Postgresql: %v\n", expected)
		}
	default:
		t.Errorf("Unknown config type: %v", confType)
	}
}

func TestLoadTarget(t *testing.T) {
	testCases := []struct {
		test     string
		name     string
		path     string
		ext      string
		expected interface{}
	}{
		{
			test: "Mongo with defaults",
			name: "mongo-non-existent",
			path: "../../config/source",
			ext:  "yaml",
			expected: target.Mongo{
				Host:        "localhost",
				Port:        27017,
				User:        "guest",
				Password:    "guest",
				Database:    "db",
				Collections: []string{},
			},
		},
		{
			test: "Mongo with custom values",
			name: "mongo",
			path: "../../config/source",
			ext:  "yaml",
			expected: target.Mongo{
				Host:        "custom-host",
				Port:        27017,
				User:        "custom-user",
				Password:    "custom-password",
				Database:    "custom-database",
				Collections: []string{"col-1", "col-2"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if strings.Contains(tc.name, "mongo") {
				config, err := LoadTarget[target.Mongo](tc.path, tc.name, tc.ext)
				Checker[target.Mongo](t, config, tc.expected, err)
			}
			if strings.Contains(tc.name, "postgresql") {
				config, err := LoadTarget[target.Postgresql](tc.path, tc.name, tc.ext)
				Checker[target.Postgresql](t, config, tc.expected, err)
			}
		})
	}
}
