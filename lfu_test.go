package lru

import (
	"testing"
	"fmt"
)

func TestSet(t *testing.T){
	cache := NewLFUCache(2)
	cache.Set("key", "Hello World")
	cache.Set("key1", "Hello go")
	cache.Set("key2", 100)
	cache.Set("key3", "Hello China")
	cache.Set("key4", 100)
	cache.Set("key4", 200)
	v,err := cache.Get("key4")
	v,err = cache.Get("key4")
	v,err = cache.Get("key4")
	if err != nil {
		t.Error("Error get data", err)
	} else {
		fmt.Println(v)
	}
	v,err = cache.Get("key3")
	v,err = cache.Get("key3")
  cache.Debug()
  cache.Set("key1", "hello lfu")
	v,err = cache.Get("key3")
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
}
func BenchmarkSet1(b *testing.B){
	for i := 0; i < b.N; i++{
		cache := New(10)
		   cache.Set("key", "Hello World")
		   cache.Set("key1", "Hello Go!")
		   cache.Set("key2", 10000)
		   cache.Set("key3", "Hello China")
		   cache.Set("key4", [...]string{"bench1", "bench2"})
		   cache.Set("key5", []string{"bench1", "bench2"})
		   v,err := cache.Get("key5")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
		   v,err = cache.Get("key4")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
		   v,err = cache.Get("key3")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
	   v,err = cache.Get("key2")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
	   v,err = cache.Get("key1")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
	   v,err = cache.Get("key")
		   if err != nil {
			   b.Error("Error get data", err)
		   } else {
			   fmt.Println(v)
		   }
	}
}
