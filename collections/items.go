package collections

type Items []interface{}

func (i *Items) Add(item interface{}) {
	newItems := append([]interface{}(*i), item)
	*i = Items(newItems)
}
