package trie_test

import "selvan.github.com/prefix"
import "testing"
import "reflect"

func assertEq(t *testing.T, first float32, second float32) {
	if first != second {
		t.Errorf("Expected %f, got %f", first, second)
	}

	t.Logf("assertEq :: First %f , second %f", first, second)	
}

func TestPut(t *testing.T) {
	_map := map[string]interface{} {
		"1" : "hsr",
		"2" : "kora",
	}

	var _trie = new(trie.Trie) 

	_trie.Put("bangalore", _map)

	_values := _trie.Get("bangalore")
	t.Log(reflect.TypeOf(_values))
	if (_values == nil) {
		t.Log(" Error !!!..Nil.")
	} else {
		t.Log(_values)
		for _key, _value := range * _values {
			t.Log(_key, _value)
		}
	}
}

func TestStartWith(t *testing.T) {
	_value := map[string]interface{} {
		"1" : "hsr",
		"2" : "kora",
	}

	_value2 := map[string]interface{} {
		"3" : "jp",
		"4" : "jaya",
	}	

	var _trie = new(trie.Trie) 

	_trie.Put("Hello", _value)
	_trie.Put("Hilly", _value)
	_trie.Put("Hello, brother", _value)
	_trie.Put("Hello, bob", _value)		

	_values := _trie.StartsWith("H", -1)
	t.Log(len(_values)==4)

	_trie.Put("Hello", _value2)
	_values = _trie.StartsWith("H", -1)
	for _key, _value := range * ((_values[0]["values"]).(*map[string]interface{})) {
		t.Log(_key, _value)
	}
}


func TestWildcard(t *testing.T) {
	_value := map[string]interface{} {
		"1" : "hsr",
		"2" : "kora",
	}

	_value2 := map[string]interface{} {
		"3" : "jp",
		"4" : "jaya",
	}	

	var _trie = new(trie.Trie) 

	_trie.Put("Hello", _value)
	_trie.Put("Hello", _value2)
	_trie.Put("Hilly", _value)
	_trie.Put("Hello, brother", _value)
	_trie.Put("Hello, bob", _value)		

	_values := _trie.Wildcard("H*...", -1)
	t.Log(len(_values))
	for _key, _value := range * ((_values[0]["values"]).(*map[string]interface{})) {
		t.Log(_key, _value)
	}
}

