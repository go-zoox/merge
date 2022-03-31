package merge

import (
	"errors"
	"reflect"
)

// Merge merges source into target.
func Merge(target interface{}, sources ...interface{}) error {
	for _, source := range sources {
		if err := merge(target, source); err != nil {
			return err
		}
	}
	return nil
}

func merge(target interface{}, source interface{}) error {
	tT := reflect.TypeOf(target)
	vS := reflect.ValueOf(source)

	// if source is nil or invalid, ignore and return
	// @TODO
	if vS.Kind() == reflect.Invalid {
		return nil
	}

	if tT.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	if vS.Kind() != reflect.Ptr {
		return errors.New("source must be a pointer")
	}

	eT := tT.Elem()

	vT := reflect.ValueOf(target).Elem()
	for i := 0; i < eT.NumField(); i++ {
		fT := eT.Field(i)

		key := fT.Name
		value := vS.Elem().FieldByName(key)

		// @TODO Nest => Should Recurrsion Deep or Shallow ?
		overwrite(vT, key, value)
	}

	return nil
}

func overwrite(vT reflect.Value, key string, value reflect.Value) {
	vT.FieldByName(key).Set(value)
}
