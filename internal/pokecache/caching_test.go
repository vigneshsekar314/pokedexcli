package pokecache

import (
	"testing"
	"time"
)

func TestCaching(t *testing.T) {
	newc := NewCache(time.Second * 1)
	testVal := "this is a sample cache value"
	testVal2 := "this is a sample cache value 2"
	newc.Add("cache1", []byte(testVal))
	newc.Add("cache2", []byte(testVal2))
	c1, ok1 := newc.Get("cache1")
	c2, ok2 := newc.Get("cache2")
	if !ok1 || !ok2 {
		t.Errorf("testfailed: value does not exist in cache")
	}
	if string(c1) != testVal || string(c2) != testVal2 {
		t.Errorf("retrieved cache values not the same as before")
	}
	time.Sleep(time.Second * 2)
	val1, nok1 := newc.Get("cache1")
	val2, nok2 := newc.Get("cache2")
	if nok1 || nok2 {
		t.Errorf("cache should not exists. Cache exists after interval - %s, %s", string(val1), string(val2))
	}

}
