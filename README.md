# popularcache
A fun generic cache that keeps the most recently added items at the front of the list and when ever an item is accessed it is pushed to the front. Order is always based on the most recently accessed / added.


```golang
    package example

    import "github.com/andrewwormald/popularcache" 

    type Item struct {
        ID    string
        Value string
    }

    func main() {
		c := popularcache.New[Item]()
		
		// Elements are added from the front
		c.Add("1", Item{ID: "1", Value: "item 1"})
		c.Add("2", Item{ID: "2", Value: "item 2"})
		c.Add("3", Item{ID: "3", Value: "item 3"})
		c.Add("4", Item{ID: "4", Value: "item 4"})

		// Calling c.List() will return:
		// []Item{
		//    {ID: "4", Value: "item 4"},
		//    {ID: "3", Value: "item 3"},
		//    {ID: "2", Value: "item 2"},
		//    {ID: "1", Value: "item 1"},
		// }

		// Calling collect moves the item to the front
		c.Collect("2")

		// Calling c.List() will return:
		// []Item{
		//    {ID: "2", Value: "item 2"},
		//    {ID: "4", Value: "item 4"},
		//    {ID: "3", Value: "item 3"},
		//    {ID: "1", Value: "item 1"},
		// }
    }
```