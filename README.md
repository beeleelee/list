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
	intList.Each(func(v lee.Item, i int){
		fmt.Println(v, i)
	})
	// lee.Each(intList, func(v lee.Item, i int){
	// 	fmt.Println(v, i)
	// })

	/*
	 *	0 0
	 *	1 1
	 *	2 2
	 */

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
}
```


## FUNCTIONS

#### func Contains(list List, f FilterFn) (r bool)
    Contains - like Find

    return true if find the item return false if can not find the item

#### func Each(list List, f EachFn)
    Each - each loop

    use for loop to get item from list and feed item to EachFn

#### func Equal(s, t List, f CmpFn) (r bool)
    Equal - a way to compare whether two list is equal

    it accept a CmpFn which handle the equal logic

#### func FindIndex(list List, f FilterFn) (index int)
    FindIndex - a way to find the index of a specific item

	it return -1 if could not find the item
	it accept a FilterFn which will specific the item

#### func SumFloat64s(l []float64) float64

#### func SumInts(l []int) int

## TYPES

#### type CmpFn func(a, b Item) bool
    CmpFn compare handle signature

    func(a, b Item) bool {

	return a == b

    }

#### type EachFn func(Item, int)
    EachFn each loop handle signature

    func(v Item, i int){

	// switch value to the expected type
	sv, _ := v.(int) // just for example, actually can use any type you specified
	fmt.Println(sv)

    }

#### type FilterFn func(Item, int) bool
    FilterFn filter loop handle signature

    func(v Item, i int) bool {

	sv := v.(string)
	return sv == "foo"

    }

#### type Item interface{}
    Item - generic type for list item

    in order to accept any type of item in collection

#### func Find(list List, f FilterFn) (r Item, ok bool)
    Find - like FindIndex, but not return index of item

    it returns the specific item and ok flag

#### func Reduce(list List, f ReduceFn, a Item) (r Item)
    Reduce - fold the list

#### type List []Item
    List a struct wrap collection in Data field

#### func Filter(list List, f FilterFn) List
    Filter - filter loop

    first create a new list then use each loop to get item from list and
    feed item to FilterFn which decide weather keep it or not

#### func From(source interface{}) (nl List, e error)
    From - convert regular slice to List

	as do not know the item type in the slic
	so use reflect package to get the item type
	and rebuild a new slice with Item type

	call it like this:
	list.From([]int{1,2,3})

#### func FromFloat64s(source []float64) (nl List)

#### func FromInts(source []int) (nl List)

#### func FromStrings(source []string) (nl List)

#### func Map(list List, f MapFn) List
    Map - map loop

    use for loop to get item from list and feed item to MapFn

#### func New(length int) List
    New generate a new List instance

#### func (l List) Contains(f FilterFn) bool
    Contains convenience wrapper for Contains function

#### func (l List) Each(f EachFn) List
    Each convenience wrapper for Each function

#### func (l List) Equal(t List, f CmpFn) bool
    Equal convenience wrapper for Equal function

#### func (l List) Filter(f FilterFn) List
    Filter convenience wrapper for Filter function

#### func (l List) Find(f FilterFn) (Item, bool)
    Find convenience wrapper for Find function

#### func (l List) FindIndex(f FilterFn) int
    FindIndex convenience wrapper for FindIndex function

#### func (l List) Map(f MapFn) List
    Map convenience wrapper for Map function

#### func (l List) Reduce(f ReduceFn, a Item) Item
    Reduce convenience wrapper for Reduce function

#### type MapFn func(Item, int) Item
    MapFn map loop handle signature

    func(v Item, i int) (item Item) {

	sv, _ := v.(float64)
	return sv * sv

    }

#### type ReduceFn func(a, b Item) Item
    ReduceFn reduce handle signature