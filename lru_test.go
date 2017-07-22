package lru

import (
	"testing"
	"fmt"
)

/* func TestSet(t *testing.T){ */
	// cache := New(10)
	// cache.Set("key", "liuyuansheng")
	// cache.Set("key1", "gengxiaoqian")
	// cache.Set("key2", "liuxiya")
	// cache.Set("key3", "hubeisheng")
	// cache.Set("key3", "hebei")
	// v,err := cache.Get("key")
	// if err != nil {
	//     t.Error("Error get data", err)
	// } else {
	//     fmt.Println(v)
	// }
/* } */
func TestSet1(t *testing.T){
	cache := New(2)
	cache.Set("key", "liuyuansheng")
	cache.Set("key1", "gengxiaoqian")
	cache.Set("key2", "liuxiya")
	cache.Set("key3", "hubeisheng")
	cache.Set("key3", "hebei")
	v,err := cache.Get("key3")
	if err != nil {
		t.Error("Error get data", err)
	} else {
		fmt.Println(v)
	}
	v,err = cache.Get("key2")
	if err != nil {
		t.Error("Error get data", err)
	} else {
		fmt.Println(v)
	}
	v,err = cache.Get("key1")
	if err != nil {
		t.Error("Error get data", err)
	} else {
		fmt.Println(v)
	}
	v,err = cache.Get("key")
	if err != nil {
		t.Error("Error get data", err)
	} else {
		fmt.Println(v)
	}
}
