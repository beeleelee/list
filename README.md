Package list

	list provide some useful utilities

	in order to manipulate collection conveniently

	in the form of functional programming

	hope it will be helpful


Example

``` package main

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

FUNCTIONS

func Contains(list List, f ItemTestFn) (r bool)
    Contains - like Find

    return true if find the item return false if can not find the item

func Each(list List, f EachFn)
    Each - each loop

    use for loop to get item from list and feed item to EachFn

func Equal(s, t List, f EqualFn) (r bool)
    Equal - a way to compare whether two list is equal

    it accept a EqualFn which handle the equal logic

func Every(list List, f ItemTestFn) (r bool)
    Every - return true if every item pass test

func FindIndex(list List, f ItemTestFn) (index int)
    FindIndex - a way to find the index of a specific item

	it return -1 if could not find the item
	it accept a ItemTestFn which will specific the item

func IsSorted(list List, f LessFn) bool
    IsSorted - convenience wrapper for std sort.SliceIsSorted

func Some(list List, f ItemTestFn) (r bool)
    Some - return true if any item pass test

func Sort(list List, f LessFn)
    Sort - convenience wrapper for std sort.Slice

TYPES

type EachFn func(Item, int)
    EachFn each loop handle signature

    func(v Item, i int){

	// switch value to the expected type
	sv, _ := v.(int) // just for example, actually can use any type you specified
	fmt.Println(sv)

    }

type EqualFn func(a, b Item) bool
    EqualFn compare handle signature

    func(a, b Item) bool {

	return a == b

    }

type Item interface{}
    Item - generic type for list item

    in order to accept any type of item in collection

func Find(list List, f ItemTestFn) (r Item, ok bool)
    Find - like FindIndex, but not return index of item

    it returns the specific item and ok flag

func Get(list List, i int) Item
    Get - get item from list it can accept negative int as index, like -1
    attention: it will never failed if then index out of range, or no item
    in list, it will return nil

func Reduce(list List, f ReduceFn, a Item) (r Item)
    Reduce - fold the list

type ItemTestFn func(Item, int) bool
    ItemTestFn filter loop handle signature

    func(v Item, i int) bool {

	sv := v.(string)
	return sv == "foo"

    }

type LessFn func(i, j int) bool
    LessFn same signature as sort.Less

type List []Item
    List a struct wrap collection in Data field

func Difference(s List, t List, f EqualFn) (r List)
    Difference - return a list with items not in the other list

func Filter(list List, f ItemTestFn) List
    Filter - filter loop

    first create a new list then use each loop to get item from list and
    feed item to ItemTestFn which decide weather keep it or not

func From(source interface{}) (nl List, e error)
    From - convert regular slice to List

	as do not know the item type in the slic
	so use reflect package to get the item type
	and rebuild a new slice with Item type

	call it like this:
	list.From([]int{1,2,3})

func FromFloat64s(source []float64) (nl List)
    FromFloat64s convert float64 slice to List

func FromInts(source []int) (nl List)
    FromInts convert int slice to List

func FromStrings(source []string) (nl List)
    FromStrings convert string slice to List

func Intersection(s List, t List, f EqualFn) (r List)
    Intersection - return a list with items in both list

func Map(list List, f MapFn) List
    Map - map loop

    use for loop to get item from list and feed item to MapFn

func New(length int) List
    New generate a new List instance

func Shuffle(list List) (r List)
    Shuffle - return a shuffled list

func Tail(list List, n int) List
    Tail - get items from last

func Union(s List, t List) List
    Union - union two lists

func (l List) Contains(f ItemTestFn) bool
    Contains convenience wrapper for Contains function

func (l List) Difference(t List, f EqualFn) List
    Difference convenience wrapper for Difference Function

func (l List) Each(f EachFn) List
    Each convenience wrapper for Each function

func (l List) Equal(t List, f EqualFn) bool
    Equal convenience wrapper for Equal function

func (l List) Every(f ItemTestFn) bool
    Every convenience wrapper for Every Function

func (l List) Filter(f ItemTestFn) List
    Filter convenience wrapper for Filter function

func (l List) Find(f ItemTestFn) (Item, bool)
    Find convenience wrapper for Find function

func (l List) FindIndex(f ItemTestFn) int
    FindIndex convenience wrapper for FindIndex function

func (l List) Get(i int) Item
    Get convenience wrapper for Get Function

func (l List) Intersection(t List, f EqualFn) List
    Intersection convenience wrapper for Intersection Function

func (l List) IsSorted(f LessFn) bool
    IsSorted convenience wrapper for std sort.SliceIsSorted

func (l List) Map(f MapFn) List
    Map convenience wrapper for Map function

func (l List) Reduce(f ReduceFn, a Item) Item
    Reduce convenience wrapper for Reduce function

func (l List) Shuffle() List
    Shuffle convenience wrapper for Shuffle Function

func (l List) Some(f ItemTestFn) bool
    Some convenience wrapper for Some Function

func (l List) Sort(f LessFn) List
    Sort convenience wrapper for std sort.Slice

func (l List) Tail(n int) List
    Tail convenience wrapper for Tail Function

func (l List) Union(t List) List
    Union convenience wrapper for Union Function

type MapFn func(Item, int) Item
    MapFn map loop handle signature

    func(v Item, i int) (item Item) {

	sv, _ := v.(float64)
	return sv * sv

    }

type ReduceFn func(a, b Item) Item
    ReduceFn reduce handle signature

type SwapFn func(i, j int)
    SwapFn swap items by index