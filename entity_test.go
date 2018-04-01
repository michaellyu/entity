package entity_test

import (
  "testing"
  "entity"
  "time"
)

var data interface{} = map[string]interface{} {
  "name": "Team",
  "city": map[string]interface{} {
    "name": "Tianjin",
  },
  "members": []interface{} {
    map[string]interface{} {
      "name": "Coral",
    },
    map[string]interface{} {
      "name": "Jack",
    },
    map[string]interface{} {
      "name": "Xinxing",
    },
  },
  "interface": 1,
  "map": map[string]interface{} { "map": 1 },
  "string": "string",
  "bool": true,
  "time": time.Now(),
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
  team := entity.NewEntity(&data)
  name, err := team.GetString("name")
  if err != nil {
    t.Error("name not found")
  }
  t.Log("name:", name)
  cityName, err := team.GetString("city.name")
  if err != nil {
    t.Error("city.name not found")
  }
  t.Log("city.name:", cityName)
  firstMember, err := team.GetMap("members[0]")
  if err != nil {
    t.Error("members[0] not found")
  }
  t.Log("members[0]:", firstMember)
  members, err := team.GetArray("members")
  if err != nil {
    t.Error("members not found")
  }
  t.Log("members:", members)
  _, err = team.Get("interface")
  if err != nil {
    t.Error("Bad Get")
  }
  _, err = team.GetField("interface")
  if err != nil {
    t.Error("Bad Field")
  }
  _, err = team.GetMap("map")
  if err != nil {
    t.Error("Bad GetMap")
  }
  _, err = team.GetString("string")
  if err != nil {
    t.Error("Bad GetString")
  }
  _, err = team.GetBool("bool")
  if err != nil {
    t.Error("Bad GetBool")
  }
  _, err = team.GetTime("time")
  if err != nil {
    t.Error("Bad GetTime")
  }
  _, err = team.GetInt("int")
  if err != nil {
    t.Error("Bad GetInt")
  }
  _, err = team.GetInt8("int8")
  if err != nil {
    t.Error("Bad GetInt8")
  }
  _, err = team.GetInt16("int16")
  if err != nil {
    t.Error("Bad GetInt16")
  }
  _, err = team.GetInt32("int32")
  if err != nil {
    t.Error("Bad GetInt32")
  }
  _, err = team.GetInt64("int64")
  if err != nil {
    t.Error("Bad GetInt64")
  }
  _, err = team.GetUint("uint")
  if err != nil {
    t.Error("Bad GetUint")
  }
  _, err = team.GetUint8("uint8")
  if err != nil {
    t.Error("Bad GetUint8")
  }
  _, err = team.GetUint16("uint16")
  if err != nil {
    t.Error("Bad GetUint16")
  }
  _, err = team.GetUint32("uint32")
  if err != nil {
    t.Error("Bad GetUint32")
  }
  _, err = team.GetUint64("uint64")
  if err != nil {
    t.Error("Bad GetUint64")
  }
  _, err = team.GetFloat32("float32")
  if err != nil {
    t.Error("Bad GetFloat32")
  }
  _, err = team.GetFloat64("float64")
  if err != nil {
    t.Error("Bad GetFloat64")
  }
  _, err = team.GetComplex64("complex64")
  if err != nil {
    t.Error("Bad GetComplex64")
  }
  _, err = team.GetComplex128("complex128")
  if err != nil {
    t.Error("Bad GetComplex128")
  }
  _, err = team.GetArray("array")
  if err != nil {
    t.Error("Bad GetArray")
  }
  _, err = team.GetArrayMap("array-map")
  if err != nil {
    t.Error("Bad GetArrayMap")
  }
  _, err = team.GetArrayBool("array-bool")
  if err != nil {
    t.Error("Bad GetArrayBool")
  }
  _, err = team.GetArrayInt("array-int")
  if err != nil {
    t.Error("Bad GetArrayInt")
  }
  _, err = team.GetArrayInt8("array-int8")
  if err != nil {
    t.Error("Bad GetArrayInt8")
  }
  _, err = team.GetArrayInt16("array-int16")
  if err != nil {
    t.Error("Bad GetArrayInt16")
  }
  _, err = team.GetArrayInt32("array-int32")
  if err != nil {
    t.Error("Bad GetArrayInt32")
  }
  _, err = team.GetArrayInt64("array-int64")
  if err != nil {
    t.Error("Bad GetArrayInt64")
  }
  _, err = team.GetArrayUint("array-uint")
  if err != nil {
    t.Error("Bad GetArrayUint")
  }
  _, err = team.GetArrayUint8("array-uint8")
  if err != nil {
    t.Error("Bad GetArrayUint8")
  }
  _, err = team.GetArrayUint16("array-uint16")
  if err != nil {
    t.Error("Bad GetArrayUint16")
  }
  _, err = team.GetArrayUint32("array-uint32")
  if err != nil {
    t.Error("Bad GetArrayUint32")
  }
  _, err = team.GetArrayUint64("array-uint64")
  if err != nil {
    t.Error("Bad GetArrayUint64")
  }
  _, err = team.GetArrayFloat32("array-float32")
  if err != nil {
    t.Error("Bad GetArrayFloat32")
  }
  _, err = team.GetArrayFloat64("array-float64")
  if err != nil {
    t.Error("Bad GetArrayFloat64")
  }
  _, err = team.GetArrayComplex64("array-complex64")
  if err != nil {
    t.Error("Bad GetArrayComplex64")
  }
  _, err = team.GetArrayComplex128("array-complex128")
  if err != nil {
    t.Error("Bad GetArrayComplex128")
  }
  _ = team.Set("name", "One Team")
  name , _ = team.GetString("name")
  if name != "One Team" {
    t.Error("Can not set field")
  }
}

func BenchmarkAll(b *testing.B) {
  team := entity.NewEntity(&data)
  for i := 0; i < b.N; i++ {
    _, _ = team.GetString("members[1].name")
  }
}
