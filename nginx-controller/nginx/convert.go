package nginx

import (
	"fmt"
	"strconv"
	"strings"

	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/runtime"
)

// There seems to be no composite interface in the kubernetes api package,
// so we have to declare our own.
type apiObject interface {
	meta.Object
	runtime.Object
}

// GetMapKeyAsBool searches the map for the given key and parses the key as bool
func GetMapKeyAsBool(m map[string]string, key string, context apiObject) (bool, bool, error) {
	if str, exists := m[key]; exists {
		b, err := strconv.ParseBool(str)
		if err != nil {
			return false, exists, fmt.Errorf("%s %v/%v '%s' contains invalid bool: %v, ignoring", context.GetObjectKind().GroupVersionKind().Kind, context.GetNamespace(), context.GetName(), key, err)
		}
		return b, exists, nil
	}
	return false, false, nil
}

// GetMapKeyAsInt tries to find and parse a key in a map as int64
func GetMapKeyAsInt(m map[string]string, key string, context apiObject) (int64, bool, error) {
	if str, exists := m[key]; exists {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, exists, fmt.Errorf("%s %v/%v '%s' contains invalid integer: %v, ignoring", context.GetObjectKind().GroupVersionKind().Kind, context.GetNamespace(), context.GetName(), key, err)
		}
		return i, exists, nil
	}
	return 0, false, nil
}

func GetMapKeyAsUint16(m map[string]string, key string, context apiObject) (uint16, bool, error) {
	if str, exists := m[key]; exists {
		i, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return 0, exists, fmt.Errorf("%s %v/%v '%s' contains invalid unsigned 16 bit integer: %v, ignoring", context.GetObjectKind().GroupVersionKind().Kind, context.GetNamespace(), context.GetName(), key, err)
		}
		return uint16(i), exists, nil
	}
	return uint16(0), false, nil
}

// GetMapKeyAsStringSlice tries to find and parse a key in the map as string slice splitting it on ','
func GetMapKeyAsStringSlice(m map[string]string, key string, context apiObject) ([]string, bool, error) {
	if str, exists := m[key]; exists {
		slice := strings.Split(str, ",")
		return slice, exists, nil
	}
	return nil, false, nil
}
