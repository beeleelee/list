# list

list provide some useful utilities 

in order to manipulate collection conveniently

in the form of functional programming 

hope, it will be helpful

## Example - basic usage

```
package main

import (
	"fmt"
	lee "github.com/beeleelee/list"
)

func main() {
	intList, _ := lee.From([]int{0,1,2})
	// list.Each
	lee.Each(&intList, func(v lee.Item, i int){
		fmt.Println(v, i)
	})
	/*
	 *	0 0
	 *	1 1
	 *	2 2
	 */

	// list.Map
	intListMapped := lee.Map(&intList, func(v lee.Item, i int) lee.Item {
		return v.(int) * 2
	})

	fmt.Println(intListMapped.Data)
	// &{[0 2 4]}

	// list.Filter
	intListFiltered := lee.Filter(&intList, func(v lee.Item, i int) bool {
		return v.(int) % 2 == 1
	})
	
	fmt.Println(intListFiltered.Data)
	// &{[1]}
}
```