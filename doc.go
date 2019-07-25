/*
Package list

	list provide some useful utilities

	in order to manipulate collection conveniently

	in the form of functional programming

	hope it will be helpful

Example

```
package main

import (
	"fmt"
	lee "github.com/beeleelee/list"
)

func main() {
	intList, _ := lee.From([]int{0,1,2})
	// list.Each
	intList.Each(func(v lee.Item, i int){
		fmt.Println(v, i)
	})
	// lee.Each(intList, func(v lee.Item, i int){
	// 	fmt.Println(v, i)
	// })

	//	0 0
	//	1 1
	//	2 2

	// list.Map
	intListMapped := intList.Map(func(v lee.Item, i int) lee.Item {
		return v.(int) * 2
	})
	// intListMapped := lee.Map(intList, func(v lee.Item, i int) lee.Item {
	// 	return v.(int) * 2
	// })

	fmt.Println(intListMapped)
	// [0 2 4]

	// list.Filter
	intListFiltered := intList.Filter(func(v lee.Item, i int) bool {
		return v.(int) % 2 == 1
	})

	// intListFiltered := lee.Filter(intList, func(v lee.Item, i int) bool {
	// 	return v.(int) % 2 == 1
	// })

	fmt.Println(intListFiltered)
	// [1]

	lee.FromInts([]int{1,2,3,4,5})
		.Reduce(func(a, b lee.Item) lee.Item {
			return a.(int) + b.(int)
		}, nil)
	// 15

	lee.FromInts([]int{3,6,9,12})
		.Intersection(leeFromInts([]{2,4,6,8,10,12}), func(a, b lee.Item) bool {
			return a == b
		})
	// [6,12]
}
```
*/
package list
