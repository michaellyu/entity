package entity_test

import (
  "testing"
  "entity"
)

var data interface{} = map[string]interface{} {
  "author": map[string]interface{} {
    "id": 1,
    "name": "Coral",
    "children": [] int { 1, 2, 3 },
  },
  "interface": 1,
  "map": map[string]interface{} { "map": 1 },
  "string": "string",
  "bool": true,
  "int": 1,
  "int8": int8(1),
  "int16": int16(1),
  "int32": int32(1),
  "int64": int64(1),
  "uint": uint(1),
  "uint8": uint8(1),
  "uint16": uint16(1),
  "uint32": uint32(1),
  "uint64": uint64(1),
  "float32": float32(1),
  "float64": float64(1),
  "complex64": complex64(1),
  "complex128": complex128(1),
  "array": []interface{} { 1 },
  "array-map": []map[string]interface{} { { "a": 1 } },
  "array-bool": []bool { true },
  "array-int": []int { 1 },
  "array-int8": []int8 { int8(1) },
  "array-int16": []int16 { int16(1) },
  "array-int32": []int32 { int32(1) },
  "array-int64": []int64 { int64(1) },
  "array-uint": []uint { uint(1) },
  "array-uint8": []uint8 { uint8(1) },
  "array-uint16": []uint16 { uint16(1) },
  "array-uint32": []uint32 { uint32(1) },
  "array-uint64": []uint64 { uint64(1) },
  "array-float32": []float32 { float32(1) },
  "array-float64": []float64 { float64(1) },
  "array-complex64": []complex64 { complex64(1) },
  "array-complex128": []complex128 { complex128(1) },
}

func TestAll(t *testing.T) {
  entity := entity.NewEntity(&data)
  authorName, err := entity.GetString("author.name")
  if err != nil {
    t.Error("author.name not found")
  }
  t.Log("author.name:", authorName)
  firstChild, err := entity.GetInt("author.children[0]")
  if err != nil {
    t.Error("author.children[0] not found")
  }
  t.Log("author.children[0]:", firstChild)
  children, err := entity.GetArrayInt("author.children")
  if err != nil {
    t.Error("author.children not found")
  }
  t.Log("author.children:", children)
  _, err = entity.Get("interface")
  if err != nil {
    t.Error("Bad Get")
  }
  _, err = entity.Field("interface")
  if err != nil {
    t.Error("Bad Field")
  }
  _, err = entity.GetMap("map")
  if err != nil {
    t.Error("Bad GetMap")
  }
  _, err = entity.GetString("string")
  if err != nil {
    t.Error("Bad GetString")
  }
  _, err = entity.GetBool("bool")
  if err != nil {
    t.Error("Bad GetBool")
  }
  _, err = entity.GetInt("int")
  if err != nil {
    t.Error("Bad GetInt")
  }
  _, err = entity.GetInt8("int8")
  if err != nil {
    t.Error("Bad GetInt8")
  }
  _, err = entity.GetInt16("int16")
  if err != nil {
    t.Error("Bad GetInt16")
  }
  _, err = entity.GetInt32("int32")
  if err != nil {
    t.Error("Bad GetInt32")
  }
  _, err = entity.GetInt64("int64")
  if err != nil {
    t.Error("Bad GetInt64")
  }
  _, err = entity.GetUint("uint")
  if err != nil {
    t.Error("Bad GetUint")
  }
  _, err = entity.GetUint8("uint8")
  if err != nil {
    t.Error("Bad GetUint8")
  }
  _, err = entity.GetUint16("uint16")
  if err != nil {
    t.Error("Bad GetUint16")
  }
  _, err = entity.GetUint32("uint32")
  if err != nil {
    t.Error("Bad GetUint32")
  }
  _, err = entity.GetUint64("uint64")
  if err != nil {
    t.Error("Bad GetUint64")
  }
  _, err = entity.GetFloat32("float32")
  if err != nil {
    t.Error("Bad GetFloat32")
  }
  _, err = entity.GetFloat64("float64")
  if err != nil {
    t.Error("Bad GetFloat64")
  }
  _, err = entity.GetComplex64("complex64")
  if err != nil {
    t.Error("Bad GetComplex64")
  }
  _, err = entity.GetComplex128("complex128")
  if err != nil {
    t.Error("Bad GetComplex128")
  }
  _, err = entity.GetArray("array")
  if err != nil {
    t.Error("Bad GetArray")
  }
  _, err = entity.GetArrayMap("array-map")
  if err != nil {
    t.Error("Bad GetArrayMap")
  }
  _, err = entity.GetArrayBool("array-bool")
  if err != nil {
    t.Error("Bad GetArrayBool")
  }
  _, err = entity.GetArrayInt("array-int")
  if err != nil {
    t.Error("Bad GetArrayInt")
  }
  _, err = entity.GetArrayInt8("array-int8")
  if err != nil {
    t.Error("Bad GetArrayInt8")
  }
  _, err = entity.GetArrayInt16("array-int16")
  if err != nil {
    t.Error("Bad GetArrayInt16")
  }
  _, err = entity.GetArrayInt32("array-int32")
  if err != nil {
    t.Error("Bad GetArrayInt32")
  }
  _, err = entity.GetArrayInt64("array-int64")
  if err != nil {
    t.Error("Bad GetArrayInt64")
  }
  _, err = entity.GetArrayUint("array-uint")
  if err != nil {
    t.Error("Bad GetArrayUint")
  }
  _, err = entity.GetArrayUint8("array-uint8")
  if err != nil {
    t.Error("Bad GetArrayUint8")
  }
  _, err = entity.GetArrayUint16("array-uint16")
  if err != nil {
    t.Error("Bad GetArrayUint16")
  }
  _, err = entity.GetArrayUint32("array-uint32")
  if err != nil {
    t.Error("Bad GetArrayUint32")
  }
  _, err = entity.GetArrayUint64("array-uint64")
  if err != nil {
    t.Error("Bad GetArrayUint64")
  }
  _, err = entity.GetArrayFloat32("array-float32")
  if err != nil {
    t.Error("Bad GetArrayFloat32")
  }
  _, err = entity.GetArrayFloat64("array-float64")
  if err != nil {
    t.Error("Bad GetArrayFloat64")
  }
  _, err = entity.GetArrayComplex64("array-complex64")
  if err != nil {
    t.Error("Bad GetArrayComplex64")
  }
  _, err = entity.GetArrayComplex128("array-complex128")
  if err != nil {
    t.Error("Bad GetArrayComplex128")
  }
}

func BenchmarkAll(b *testing.B) {
  entity := entity.NewEntity(&data)
  for i := 0; i < b.N; i++ {
    _, _ = entity.GetString("comments[1].author.name")
  }
}
