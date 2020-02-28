// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_kube-config", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("kube-config"); err == nil {
				assert.Equal(t, string(defaultConfig.KubeConfigPath), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("kube-config", testValue)
			if vString, err := cmdFlags.GetString("kube-config"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.KubeConfigPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_master", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("master"); err == nil {
				assert.Equal(t, string(defaultConfig.MasterURL), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("master", testValue)
			if vString, err := cmdFlags.GetString("master"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MasterURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_workers", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				assert.Equal(t, int(defaultConfig.Workers), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("workers", testValue)
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_workflow-reeval-duration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("workflow-reeval-duration"); err == nil {
				assert.Equal(t, string(defaultConfig.WorkflowReEval.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.WorkflowReEval.String()

			cmdFlags.Set("workflow-reeval-duration", testValue)
			if vString, err := cmdFlags.GetString("workflow-reeval-duration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.WorkflowReEval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_downstream-eval-duration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("downstream-eval-duration"); err == nil {
				assert.Equal(t, string(defaultConfig.DownstreamEval.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.DownstreamEval.String()

			cmdFlags.Set("downstream-eval-duration", testValue)
			if vString, err := cmdFlags.GetString("downstream-eval-duration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.DownstreamEval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_limit-namespace", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("limit-namespace"); err == nil {
				assert.Equal(t, string(defaultConfig.LimitNamespace), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("limit-namespace", testValue)
			if vString, err := cmdFlags.GetString("limit-namespace"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LimitNamespace)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_prof-port", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("prof-port"); err == nil {
				assert.Equal(t, string(defaultConfig.ProfilerPort.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.ProfilerPort.String()

			cmdFlags.Set("prof-port", testValue)
			if vString, err := cmdFlags.GetString("prof-port"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ProfilerPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metadata-prefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("metadata-prefix"); err == nil {
				assert.Equal(t, string(defaultConfig.MetadataPrefix), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metadata-prefix", testValue)
			if vString, err := cmdFlags.GetString("metadata-prefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MetadataPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.type", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.type"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Type), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.type", testValue)
			if vString, err := cmdFlags.GetString("queue.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.queue.type", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.queue.type"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Queue.Type), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.queue.type", testValue)
			if vString, err := cmdFlags.GetString("queue.queue.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Queue.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.queue.base-delay", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.queue.base-delay"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Queue.BaseDelay.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Queue.Queue.BaseDelay.String()

			cmdFlags.Set("queue.queue.base-delay", testValue)
			if vString, err := cmdFlags.GetString("queue.queue.base-delay"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Queue.BaseDelay)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.queue.max-delay", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.queue.max-delay"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Queue.MaxDelay.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Queue.Queue.MaxDelay.String()

			cmdFlags.Set("queue.queue.max-delay", testValue)
			if vString, err := cmdFlags.GetString("queue.queue.max-delay"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Queue.MaxDelay)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.queue.rate", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt64, err := cmdFlags.GetInt64("queue.queue.rate"); err == nil {
				assert.Equal(t, int64(defaultConfig.Queue.Queue.Rate), vInt64)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.queue.rate", testValue)
			if vInt64, err := cmdFlags.GetInt64("queue.queue.rate"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.Queue.Queue.Rate)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.queue.capacity", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("queue.queue.capacity"); err == nil {
				assert.Equal(t, int(defaultConfig.Queue.Queue.Capacity), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.queue.capacity", testValue)
			if vInt, err := cmdFlags.GetInt("queue.queue.capacity"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Queue.Queue.Capacity)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.sub-queue.type", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.sub-queue.type"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Sub.Type), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.sub-queue.type", testValue)
			if vString, err := cmdFlags.GetString("queue.sub-queue.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Sub.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.sub-queue.base-delay", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.sub-queue.base-delay"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Sub.BaseDelay.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Queue.Sub.BaseDelay.String()

			cmdFlags.Set("queue.sub-queue.base-delay", testValue)
			if vString, err := cmdFlags.GetString("queue.sub-queue.base-delay"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Sub.BaseDelay)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.sub-queue.max-delay", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.sub-queue.max-delay"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.Sub.MaxDelay.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Queue.Sub.MaxDelay.String()

			cmdFlags.Set("queue.sub-queue.max-delay", testValue)
			if vString, err := cmdFlags.GetString("queue.sub-queue.max-delay"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.Sub.MaxDelay)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.sub-queue.rate", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt64, err := cmdFlags.GetInt64("queue.sub-queue.rate"); err == nil {
				assert.Equal(t, int64(defaultConfig.Queue.Sub.Rate), vInt64)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.sub-queue.rate", testValue)
			if vInt64, err := cmdFlags.GetInt64("queue.sub-queue.rate"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.Queue.Sub.Rate)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.sub-queue.capacity", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("queue.sub-queue.capacity"); err == nil {
				assert.Equal(t, int(defaultConfig.Queue.Sub.Capacity), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.sub-queue.capacity", testValue)
			if vInt, err := cmdFlags.GetInt("queue.sub-queue.capacity"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Queue.Sub.Capacity)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.batching-interval", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("queue.batching-interval"); err == nil {
				assert.Equal(t, string(defaultConfig.Queue.BatchingInterval.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Queue.BatchingInterval.String()

			cmdFlags.Set("queue.batching-interval", testValue)
			if vString, err := cmdFlags.GetString("queue.batching-interval"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Queue.BatchingInterval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_queue.batch-size", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("queue.batch-size"); err == nil {
				assert.Equal(t, int(defaultConfig.Queue.BatchSize), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("queue.batch-size", testValue)
			if vInt, err := cmdFlags.GetInt("queue.batch-size"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Queue.BatchSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metrics-prefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("metrics-prefix"); err == nil {
				assert.Equal(t, string(defaultConfig.MetricsPrefix), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metrics-prefix", testValue)
			if vString, err := cmdFlags.GetString("metrics-prefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MetricsPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_enable-admin-launcher", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("enable-admin-launcher"); err == nil {
				assert.Equal(t, bool(defaultConfig.EnableAdminLauncher), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("enable-admin-launcher", testValue)
			if vBool, err := cmdFlags.GetBool("enable-admin-launcher"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.EnableAdminLauncher)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_max-workflow-retries", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("max-workflow-retries"); err == nil {
				assert.Equal(t, int(defaultConfig.MaxWorkflowRetries), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("max-workflow-retries", testValue)
			if vInt, err := cmdFlags.GetInt("max-workflow-retries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.MaxWorkflowRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_max-ttl-hours", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("max-ttl-hours"); err == nil {
				assert.Equal(t, int(defaultConfig.MaxTTLInHours), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("max-ttl-hours", testValue)
			if vInt, err := cmdFlags.GetInt("max-ttl-hours"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.MaxTTLInHours)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_gc-interval", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("gc-interval"); err == nil {
				assert.Equal(t, string(defaultConfig.GCInterval.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.GCInterval.String()

			cmdFlags.Set("gc-interval", testValue)
			if vString, err := cmdFlags.GetString("gc-interval"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.GCInterval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.enabled", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("leader-election.enabled"); err == nil {
				assert.Equal(t, bool(defaultConfig.LeaderElection.Enabled), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("leader-election.enabled", testValue)
			if vBool, err := cmdFlags.GetBool("leader-election.enabled"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.LeaderElection.Enabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.lock-config-map.Namespace", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("leader-election.lock-config-map.Namespace"); err == nil {
				assert.Equal(t, string(defaultConfig.LeaderElection.LockConfigMap.Namespace), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("leader-election.lock-config-map.Namespace", testValue)
			if vString, err := cmdFlags.GetString("leader-election.lock-config-map.Namespace"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LeaderElection.LockConfigMap.Namespace)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.lock-config-map.Name", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("leader-election.lock-config-map.Name"); err == nil {
				assert.Equal(t, string(defaultConfig.LeaderElection.LockConfigMap.Name), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("leader-election.lock-config-map.Name", testValue)
			if vString, err := cmdFlags.GetString("leader-election.lock-config-map.Name"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LeaderElection.LockConfigMap.Name)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.lease-duration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("leader-election.lease-duration"); err == nil {
				assert.Equal(t, string(defaultConfig.LeaderElection.LeaseDuration.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.LeaderElection.LeaseDuration.String()

			cmdFlags.Set("leader-election.lease-duration", testValue)
			if vString, err := cmdFlags.GetString("leader-election.lease-duration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LeaderElection.LeaseDuration)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.renew-deadline", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("leader-election.renew-deadline"); err == nil {
				assert.Equal(t, string(defaultConfig.LeaderElection.RenewDeadline.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.LeaderElection.RenewDeadline.String()

			cmdFlags.Set("leader-election.renew-deadline", testValue)
			if vString, err := cmdFlags.GetString("leader-election.renew-deadline"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LeaderElection.RenewDeadline)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_leader-election.retry-period", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("leader-election.retry-period"); err == nil {
				assert.Equal(t, string(defaultConfig.LeaderElection.RetryPeriod.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.LeaderElection.RetryPeriod.String()

			cmdFlags.Set("leader-election.retry-period", testValue)
			if vString, err := cmdFlags.GetString("leader-election.retry-period"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LeaderElection.RetryPeriod)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_publish-k8s-events", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("publish-k8s-events"); err == nil {
				assert.Equal(t, bool(defaultConfig.PublishK8sEvents), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("publish-k8s-events", testValue)
			if vBool, err := cmdFlags.GetBool("publish-k8s-events"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.PublishK8sEvents)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_max-output-size-bytes", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt64, err := cmdFlags.GetInt64("max-output-size-bytes"); err == nil {
				assert.Equal(t, int64(defaultConfig.MaxDatasetSizeBytes), vInt64)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("max-output-size-bytes", testValue)
			if vInt64, err := cmdFlags.GetInt64("max-output-size-bytes"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.MaxDatasetSizeBytes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_kube-client-config.burst", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("kube-client-config.burst"); err == nil {
				assert.Equal(t, int(defaultConfig.KubeConfig.Burst), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("kube-client-config.burst", testValue)
			if vInt, err := cmdFlags.GetInt("kube-client-config.burst"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.KubeConfig.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_kube-client-config.timeout", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("kube-client-config.timeout"); err == nil {
				assert.Equal(t, string(defaultConfig.KubeConfig.Timeout.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.KubeConfig.Timeout.String()

			cmdFlags.Set("kube-client-config.timeout", testValue)
			if vString, err := cmdFlags.GetString("kube-client-config.timeout"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.KubeConfig.Timeout)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-deadlines.node-execution-deadline", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("default-deadlines.node-execution-deadline"); err == nil {
				assert.Equal(t, string(defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeExecutionDeadline.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeExecutionDeadline.String()

			cmdFlags.Set("default-deadlines.node-execution-deadline", testValue)
			if vString, err := cmdFlags.GetString("default-deadlines.node-execution-deadline"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.NodeConfig.DefaultDeadlines.DefaultNodeExecutionDeadline)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-deadlines.node-active-deadline", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("default-deadlines.node-active-deadline"); err == nil {
				assert.Equal(t, string(defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeActiveDeadline.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeActiveDeadline.String()

			cmdFlags.Set("default-deadlines.node-active-deadline", testValue)
			if vString, err := cmdFlags.GetString("default-deadlines.node-active-deadline"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.NodeConfig.DefaultDeadlines.DefaultNodeActiveDeadline)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-deadlines.workflow-active-deadline", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("default-deadlines.workflow-active-deadline"); err == nil {
				assert.Equal(t, string(defaultConfig.NodeConfig.DefaultDeadlines.DefaultWorkflowActiveDeadline.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.NodeConfig.DefaultDeadlines.DefaultWorkflowActiveDeadline.String()

			cmdFlags.Set("default-deadlines.workflow-active-deadline", testValue)
			if vString, err := cmdFlags.GetString("default-deadlines.workflow-active-deadline"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.NodeConfig.DefaultDeadlines.DefaultWorkflowActiveDeadline)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
