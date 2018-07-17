package entity

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

type Entity struct {
	entity *interface{}
}

func NewEntity(data *interface{}) Entity {
	return Entity{
		entity: data,
	}
}

func (self Entity) Get(name string) (interface{}, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case interface{}:
		return field.(interface{}), nil
	default:
		return nil, errors.New("Cannot convert to an interface{}!")
	}
}

func (self Entity) GetMap(name string) (map[string]interface{}, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case map[string]interface{}:
		return field.(map[string]interface{}), nil
	default:
		return nil, errors.New("Cannot convert to a map[string]interface{}!")
	}
}

func (self Entity) GetString(name string) (string, error) {
	field, err := self.GetField(name)
	if err != nil {
		return "", err
	}
	switch field.(type) {
	case string:
		return field.(string), nil
	default:
		return "", errors.New("Cannot convert to a string!")
	}
}

func (self Entity) GetBool(name string) (bool, error) {
	field, err := self.GetField(name)
	if err != nil {
		return false, err
	}
	switch field.(type) {
	case bool:
		return field.(bool), nil
	default:
		return false, errors.New("Cannot convert to a bool!")
	}
}

func (self Entity) GetTime(name string) (time.Time, error) {
	field, err := self.GetField(name)
	if err != nil {
		return time.Time{}, err
	}
	switch field.(type) {
	case time.Time:
		return field.(time.Time), nil
	default:
		return time.Time{}, errors.New("Cannot convert to a Time!")
	}
}

func (self Entity) GetInt(name string) (int, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case int:
		return field.(int), nil
	default:
		return 0, errors.New("Cannot convert to an int!")
	}
}

func (self Entity) GetInt8(name string) (int8, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case int8:
		return field.(int8), nil
	default:
		return 0, errors.New("Cannot convert to an int8!")
	}
}

func (self Entity) GetInt16(name string) (int16, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case int16:
		return field.(int16), nil
	default:
		return 0, errors.New("Cannot convert to an int16!")
	}
}

func (self Entity) GetInt32(name string) (int32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case int32:
		return field.(int32), nil
	default:
		return 0, errors.New("Cannot convert to an int32!")
	}
}

func (self Entity) GetInt64(name string) (int64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case int64:
		return field.(int64), nil
	default:
		return 0, errors.New("Cannot convert to an int64!")
	}
}

func (self Entity) GetUint(name string) (uint, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case uint:
		return field.(uint), nil
	default:
		return 0, errors.New("Cannot convert to an uint!")
	}
}

func (self Entity) GetUint8(name string) (uint8, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case uint8:
		return field.(uint8), nil
	default:
		return 0, errors.New("Cannot convert to an uint8!")
	}
}

func (self Entity) GetUint16(name string) (uint16, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case uint16:
		return field.(uint16), nil
	default:
		return 0, errors.New("Cannot convert to an uint16!")
	}
}

func (self Entity) GetUint32(name string) (uint32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case uint32:
		return field.(uint32), nil
	default:
		return 0, errors.New("Cannot convert to an uint32!")
	}
}

func (self Entity) GetUint64(name string) (uint64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case uint64:
		return field.(uint64), nil
	default:
		return 0, errors.New("Cannot convert to an uint64!")
	}
}

func (self Entity) GetFloat32(name string) (float32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case float32:
		return field.(float32), nil
	default:
		return 0, errors.New("Cannot convert to a float32!")
	}
}

func (self Entity) GetFloat64(name string) (float64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case float64:
		return field.(float64), nil
	default:
		return 0, errors.New("Cannot convert to a float64!")
	}
}

func (self Entity) GetComplex64(name string) (complex64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case complex64:
		return field.(complex64), nil
	default:
		return 0, errors.New("Cannot convert to a complex64!")
	}
}

func (self Entity) GetComplex128(name string) (complex128, error) {
	field, err := self.GetField(name)
	if err != nil {
		return 0, err
	}
	switch field.(type) {
	case complex128:
		return field.(complex128), nil
	default:
		return 0, errors.New("Cannot convert to a complex128!")
	}
}

func (self Entity) GetArray(name string) ([]interface{}, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []interface{}:
		return field.([]interface{}), nil
	default:
		return nil, errors.New("Cannot convert to a []interface{}!")
	}
}

func (self Entity) GetArrayMap(name string) ([]map[string]interface{}, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []map[string]interface{}:
		return field.([]map[string]interface{}), nil
	default:
		return nil, errors.New("Cannot convert to a []map[string]interface{}!")
	}
}

func (self Entity) GetArrayBool(name string) ([]bool, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []bool:
		return field.([]bool), nil
	default:
		return nil, errors.New("Cannot convert to a []bool!")
	}
}

func (self Entity) GetArrayInt(name string) ([]int, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []int:
		return field.([]int), nil
	default:
		return nil, errors.New("Cannot convert to a []int!")
	}
}

func (self Entity) GetArrayInt8(name string) ([]int8, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []int8:
		return field.([]int8), nil
	default:
		return nil, errors.New("Cannot convert to a []int8!")
	}
}

func (self Entity) GetArrayInt16(name string) ([]int16, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []int16:
		return field.([]int16), nil
	default:
		return nil, errors.New("Cannot convert to a []int16!")
	}
}

func (self Entity) GetArrayInt32(name string) ([]int32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []int32:
		return field.([]int32), nil
	default:
		return nil, errors.New("Cannot convert to a []int32!")
	}
}

func (self Entity) GetArrayInt64(name string) ([]int64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []int64:
		return field.([]int64), nil
	default:
		return nil, errors.New("Cannot convert to a []int64!")
	}
}

func (self Entity) GetArrayUint(name string) ([]uint, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []uint:
		return field.([]uint), nil
	default:
		return nil, errors.New("Cannot convert to a []uint!")
	}
}

func (self Entity) GetArrayUint8(name string) ([]uint8, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []uint8:
		return field.([]uint8), nil
	default:
		return nil, errors.New("Cannot convert to a []uint8!")
	}
}

func (self Entity) GetArrayUint16(name string) ([]uint16, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []uint16:
		return field.([]uint16), nil
	default:
		return nil, errors.New("Cannot convert to a []uint16!")
	}
}

func (self Entity) GetArrayUint32(name string) ([]uint32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []uint32:
		return field.([]uint32), nil
	default:
		return nil, errors.New("Cannot convert to a []uint32!")
	}
}

func (self Entity) GetArrayUint64(name string) ([]uint64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []uint64:
		return field.([]uint64), nil
	default:
		return nil, errors.New("Cannot convert to a []uint64!")
	}
}

func (self Entity) GetArrayFloat32(name string) ([]float32, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []float32:
		return field.([]float32), nil
	default:
		return nil, errors.New("Cannot convert to a []float32!")
	}
}

func (self Entity) GetArrayFloat64(name string) ([]float64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []float64:
		return field.([]float64), nil
	default:
		return nil, errors.New("Cannot convert to a []float64!")
	}
}

func (self Entity) GetArrayComplex64(name string) ([]complex64, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []complex64:
		return field.([]complex64), nil
	default:
		return nil, errors.New("Cannot convert to a []complex64!")
	}
}

func (self Entity) GetArrayComplex128(name string) ([]complex128, error) {
	field, err := self.GetField(name)
	if err != nil {
		return nil, err
	}
	switch field.(type) {
	case []complex128:
		return field.([]complex128), nil
	default:
		return nil, errors.New("Cannot convert to a []complex128!")
	}
}

var rIndex *regexp.Regexp = regexp.MustCompile(`^\[([\d]+)\]$`)
var rIndexWrapper *regexp.Regexp = regexp.MustCompile(`[\[\]]`)

func (self Entity) GetField(name string) (interface{}, error) {
	if name == "" {
		return nil, errors.New("Bad field name!")
	}
	paths := nameToPaths(name)
	var current interface{} = *self.entity
	var ok bool
	for level := 0; level < len(paths); level++ {
		if rIndex.MatchString(paths[level]) {
			index, _ := strconv.Atoi(rIndexWrapper.ReplaceAllString(paths[level], ""))
			switch current.(type) {
			case []interface{}:
				tmp := current.([]interface{})
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []string:
				tmp := current.([]string)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []bool:
				tmp := current.([]bool)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []time.Time:
				tmp := current.([]time.Time)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []int:
				tmp := current.([]int)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []int8:
				tmp := current.([]int8)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []int16:
				tmp := current.([]int16)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []int32:
				tmp := current.([]int32)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []int64:
				tmp := current.([]int64)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []uint:
				tmp := current.([]uint)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []uint8:
				tmp := current.([]uint8)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []uint16:
				tmp := current.([]uint16)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []uint32:
				tmp := current.([]uint32)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []uint64:
				tmp := current.([]uint64)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []float32:
				tmp := current.([]float32)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []float64:
				tmp := current.([]float64)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []complex64:
				tmp := current.([]complex64)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			case []complex128:
				tmp := current.([]complex128)
				if index >= 0 && index < len(tmp) {
					current = tmp[index]
				} else {
					current = nil
				}
			}
			if current == nil {
				return nil, errors.New("Index out of range!")
			}
		} else {
			switch current.(type) {
			case map[string]interface{}:
				current, ok = current.(map[string]interface{})[paths[level]]
			case map[string]string:
				current, ok = current.(map[string]string)[paths[level]]
			case map[string]bool:
				current, ok = current.(map[string]bool)[paths[level]]
			case map[string]time.Time:
				current, ok = current.(map[string]time.Time)[paths[level]]
			case map[string]int:
				current, ok = current.(map[string]int)[paths[level]]
			case map[string]int8:
				current, ok = current.(map[string]int8)[paths[level]]
			case map[string]int16:
				current, ok = current.(map[string]int16)[paths[level]]
			case map[string]int32:
				current, ok = current.(map[string]int32)[paths[level]]
			case map[string]int64:
				current, ok = current.(map[string]int64)[paths[level]]
			case map[string]uint:
				current, ok = current.(map[string]uint)[paths[level]]
			case map[string]uint8:
				current, ok = current.(map[string]uint8)[paths[level]]
			case map[string]uint16:
				current, ok = current.(map[string]uint16)[paths[level]]
			case map[string]uint32:
				current, ok = current.(map[string]uint32)[paths[level]]
			case map[string]uint64:
				current, ok = current.(map[string]uint64)[paths[level]]
			case map[string]float32:
				current, ok = current.(map[string]float32)[paths[level]]
			case map[string]float64:
				current, ok = current.(map[string]float64)[paths[level]]
			case map[string]complex64:
				current, ok = current.(map[string]complex64)[paths[level]]
			case map[string]complex128:
				current, ok = current.(map[string]complex128)[paths[level]]
			}
			if !ok {
				return nil, errors.New("Cannot exist field!")
			}
		}
	}
	return current, nil
}

func (self Entity) Set(name string, value interface{}) error {
	return self.SetField(name, value)
}

func (self Entity) SetField(name string, value interface{}) error {
	if name == "" {
		return errors.New("Bad field name!")
	}
	paths := nameToPaths(name)
	var current interface{} = *self.entity
	var length = len(paths)
	var last = length - 1
	for level := 0; level < length; level++ {
		if rIndex.MatchString(paths[level]) {
			index, _ := strconv.Atoi(rIndexWrapper.ReplaceAllString(paths[level], ""))
			switch current.(type) {
			case []interface{}:
				tmp := current.([]interface{})
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = value
					}
				} else {
					current = nil
				}
			case []string:
				tmp := current.([]string)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = string(value.(string))
					}
				} else {
					current = nil
				}
			case []bool:
				tmp := current.([]bool)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = bool(value.(bool))
					}
				} else {
					current = nil
				}
			case []time.Time:
				tmp := current.([]time.Time)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = time.Time(value.(time.Time))
					}
				} else {
					current = nil
				}
			case []int:
				tmp := current.([]int)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = int(value.(int))
					}
				} else {
					current = nil
				}
			case []int8:
				tmp := current.([]int8)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = int8(value.(int8))
					}
				} else {
					current = nil
				}
			case []int16:
				tmp := current.([]int16)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = int16(value.(int16))
					}
				} else {
					current = nil
				}
			case []int32:
				tmp := current.([]int32)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = int32(value.(int32))
					}
				} else {
					current = nil
				}
			case []int64:
				tmp := current.([]int64)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = int64(value.(int64))
					}
				} else {
					current = nil
				}
			case []uint:
				tmp := current.([]uint)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = uint(value.(uint))
					}
				} else {
					current = nil
				}
			case []uint8:
				tmp := current.([]uint8)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = uint8(value.(uint8))
					}
				} else {
					current = nil
				}
			case []uint16:
				tmp := current.([]uint16)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = uint16(value.(uint16))
					}
				} else {
					current = nil
				}
			case []uint32:
				tmp := current.([]uint32)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = uint32(value.(uint32))
					}
				} else {
					current = nil
				}
			case []uint64:
				tmp := current.([]uint64)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = uint64(value.(uint64))
					}
				} else {
					current = nil
				}
			case []float32:
				tmp := current.([]float32)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = float32(value.(float32))
					}
				} else {
					current = nil
				}
			case []float64:
				tmp := current.([]float64)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = float64(value.(float64))
					}
				} else {
					current = nil
				}
			case []complex64:
				tmp := current.([]complex64)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = complex64(value.(complex64))
					}
				} else {
					current = nil
				}
			case []complex128:
				tmp := current.([]complex128)
				if index >= 0 && index < len(tmp) {
					if level != last {
						current = tmp[index]
					} else {
						tmp[index] = complex128(value.(complex128))
					}
				} else {
					current = nil
				}
			}
			if current == nil {
				return errors.New("Index out of range!")
			}
		} else {
			switch current.(type) {
			case map[string]interface{}:
				if level != last {
					current, _ = current.(map[string]interface{})[paths[level]]
				} else {
					current.(map[string]interface{})[paths[level]] = value
				}
			case map[string]string:
				if level != last {
					current, _ = current.(map[string]string)[paths[level]]
				} else {
					current.(map[string]string)[paths[level]] = string(value.(string))
				}
			case map[string]bool:
				if level != last {
					current, _ = current.(map[string]bool)[paths[level]]
				} else {
					current.(map[string]bool)[paths[level]] = bool(value.(bool))
				}
			case map[string]time.Time:
				if level != last {
					current, _ = current.(map[string]time.Time)[paths[level]]
				} else {
					current.(map[string]time.Time)[paths[level]] = time.Time(value.(time.Time))
				}
			case map[string]int:
				if level != last {
					current, _ = current.(map[string]int)[paths[level]]
				} else {
					current.(map[string]int)[paths[level]] = int(value.(int))
				}
			case map[string]int8:
				if level != last {
					current, _ = current.(map[string]int8)[paths[level]]
				} else {
					current.(map[string]int8)[paths[level]] = int8(value.(int8))
				}
			case map[string]int16:
				if level != last {
					current, _ = current.(map[string]int16)[paths[level]]
				} else {
					current.(map[string]int16)[paths[level]] = int16(value.(int16))
				}
			case map[string]int32:
				if level != last {
					current, _ = current.(map[string]int32)[paths[level]]
				} else {
					current.(map[string]int32)[paths[level]] = int32(value.(int32))
				}
			case map[string]int64:
				if level != last {
					current, _ = current.(map[string]int64)[paths[level]]
				} else {
					current.(map[string]int64)[paths[level]] = int64(value.(int64))
				}
			case map[string]uint:
				if level != last {
					current, _ = current.(map[string]uint)[paths[level]]
				} else {
					current.(map[string]uint)[paths[level]] = uint(value.(uint))
				}
			case map[string]uint8:
				if level != last {
					current, _ = current.(map[string]uint8)[paths[level]]
				} else {
					current.(map[string]uint8)[paths[level]] = uint8(value.(uint8))
				}
			case map[string]uint16:
				if level != last {
					current, _ = current.(map[string]uint16)[paths[level]]
				} else {
					current.(map[string]uint16)[paths[level]] = uint16(value.(uint16))
				}
			case map[string]uint32:
				if level != last {
					current, _ = current.(map[string]uint32)[paths[level]]
				} else {
					current.(map[string]uint32)[paths[level]] = uint32(value.(uint32))
				}
			case map[string]uint64:
				if level != last {
					current, _ = current.(map[string]uint64)[paths[level]]
				} else {
					current.(map[string]uint64)[paths[level]] = uint64(value.(uint64))
				}
			case map[string]float32:
				if level != last {
					current, _ = current.(map[string]float32)[paths[level]]
				} else {
					current.(map[string]float32)[paths[level]] = float32(value.(float32))
				}
			case map[string]float64:
				if level != last {
					current, _ = current.(map[string]float64)[paths[level]]
				} else {
					current.(map[string]float64)[paths[level]] = float64(value.(float64))
				}
			case map[string]complex64:
				if level != last {
					current, _ = current.(map[string]complex64)[paths[level]]
				} else {
					current.(map[string]complex64)[paths[level]] = complex64(value.(complex64))
				}
			case map[string]complex128:
				if level != last {
					current, _ = current.(map[string]complex128)[paths[level]]
				} else {
					current.(map[string]complex128)[paths[level]] = complex128(value.(complex128))
				}
			}
		}
	}
	if current == nil {
		return errors.New("Cannot exist field!")
	}
	return nil
}

var rPath *regexp.Regexp = regexp.MustCompile(`[^.\[\]]+|\[[\d]+\]`)

func nameToPaths(name string) []string {
	paths := rPath.FindAllString(name, -1)
	return paths
}
