package trie

import "strings"

type Trie struct {
	Root *Node
}

type Node struct {
	Char rune
	Values map[string]interface{}
	End bool
	Left, Right, Middle *Node
}

func (trie *Trie) Put(key string, values map[string]interface{}) {
	if key == "" {
		return
	}

	trie.Root = trie.PutRecursive(trie.Root, []rune(key), 0, &values)
}

func (trie *Trie) PutRecursive(node *Node, key []rune, index int, values *map[string]interface{}) *Node {

	char := key[index]

	if node == nil {
		node = &Node{Char : char, Values : make(map[string]interface{})}
	}

	switch {
		case char < node.Char:
			node.Left = trie.PutRecursive(node.Left, key, index, values)
		case char > node.Char:
			node.Right = trie.PutRecursive(node.Right, key, index, values)
		case index < len(key)-1 : /* We're not at the end of the input string; add next chr */
			node.Middle = trie.PutRecursive(node.Middle, key, index + 1, values)						
		default:
			node.End = true
			for _key, _value := range *values {
				node.Values[_key] = _value;
			}
	}

	return node
}

func (trie *Trie) Get(key string) *map[string]interface{} {
	return trie.GetRecursive(trie.Root, []rune(key), 0)
}

func (trie *Trie) GetRecursive(node *Node, key []rune, index int) *map[string]interface{} {
	if node == nil {
       return nil
    }

    char := key[index]

    switch {
	    case char < node.Char:
	        return trie.GetRecursive(node.Left, key, index)
	    case char > node.Char:
	        return trie.GetRecursive(node.Right, key, index)
	    case index < len(key) - 1 : /* We're not at the end of the input string; add next char */
	        return trie.GetRecursive(node.Middle, key, index + 1)
	    default:
	    	if node.End {
	    		return &node.Values
	    	} else {
	    		return nil
	    	}
    }
}

func (trie *Trie) StartsWith(key string, items_to_fetch int) []map[string]interface{} {
	var results []map[string]interface{}
	
	config := map[string]int {
		"items_to_fetch": items_to_fetch, 
		"start_count":0,
	}

	trie.StartsWithRecursive(trie.Root, []rune(key), 0, []rune(""), &results, &config);
	return results
}

func (trie *Trie) StartsWithRecursive(node *Node, key []rune, index int, prefix []rune, results *[]map[string]interface{}, config *map[string]int) {
	if node == nil || (*config)["start_count"] == (*config)["items_to_fetch"] {
        return
    }

    var char rune

    if index < len(key) {
    	char = key[index]
    } else {
    	char = rune('*')
    }

    if node.End {
    	matched_key := append(prefix, node.Char)
    	str_matched_key := string(matched_key)
    	str_key := string(key)
        if (strings.Index(str_matched_key, str_key) == 0 ) {
            *results = append(*results, map[string]interface{} {"key":str_matched_key, "values":&node.Values,});
            (*config)["start_count"] = (*config)["start_count"] + 1;
        }	
    }

	if (char == '*' || char == '.' || char < node.Char) {
        trie.StartsWithRecursive(node.Left, key, index, prefix, results, config);
    }

    if (char == '*' || char == '.' || char == node.Char) {
        trie.StartsWithRecursive(node.Middle, key, index + 1, append(prefix, node.Char), results, config);
    }

    if (char == '*' || char == '.' || char > node.Char) {
        trie.StartsWithRecursive(node.Right, key, index, prefix, results, config);
    }
}

func (trie *Trie) Wildcard(key string, items_to_fetch int) []map[string]interface{} {
	var results []map[string]interface{}
	
	config := map[string]int {
		"items_to_fetch": items_to_fetch, 
		"start_count":0,
	}

	trie.WildRecursive(trie.Root, []rune(key), 0, []rune(""), &results, &config);
	return results
}

func (trie *Trie) WildRecursive(node *Node, key []rune, index int, prefix []rune, results *[]map[string]interface{}, config *map[string]int) {
	if node == nil || index == len(key) || (*config)["start_count"] == (*config)["items_to_fetch"] {
        return
    }

    char := key[index]

    if node.End {
    	matched_key := append(prefix, node.Char)
    	str_matched_key := string(matched_key)
        if (len(matched_key) >= len(key)) {
            *results = append(*results, map[string]interface{} {"key":str_matched_key, "values":&node.Values,});
            (*config)["start_count"] = (*config)["start_count"] + 1;
        }	
    }

	if (char == '*' || char == '.' || char < node.Char) {
        trie.WildRecursive(node.Left, key, index, prefix, results, config);
    }

    if (char == '*' || char == '.' || char == node.Char) {
        trie.WildRecursive(node.Middle, key, index + 1, append(prefix, node.Char), results, config);
    }

    if (char == '*' || char == '.' || char > node.Char) {
        trie.WildRecursive(node.Right, key, index, prefix, results, config);
    }
}
