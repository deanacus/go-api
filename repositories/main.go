package repositories

import (
	// "encoding/json"
	"encoding/json"
	"errors"
	"log"
	"time"

	// "log"
	"sync"
)

// type Item interface {
// 	GetID() int
// 	SetID(id int)
// }

type WithID interface {
	GetID() int
	SetID(int)
	SetCreatedAt(time.Time)
	GetCreatedAt() time.Time
}

type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (b Base) GetID() int {
	return b.ID
}

func (b Base) SetID(id int) {
	b.ID = id
}

func (b Base) GetCreatedAt() time.Time {
	return b.CreatedAt
}

func (b Base) SetCreatedAt(createdAt time.Time) {
	b.CreatedAt = createdAt
}

type Collection[T WithID] struct {
	Items     map[int]T
	CurrentID int
	Mutex     sync.Mutex
}

func NewCollection[T WithID]() Collection[T] {
	return Collection[T]{
		Items:     make(map[int]T),
		CurrentID: 0,
		Mutex:     sync.Mutex{},
	}
}

func (coll *Collection[T]) GetOne(id int) (T, error) {
	coll.Mutex.Lock()
	defer coll.Mutex.Unlock()
	item, ok := coll.Items[id]
	if !ok {
		return item, errors.New("not found")
	}
	return item, nil
}

func (coll *Collection[T]) GetMany() []T {
	coll.Mutex.Lock()
	defer coll.Mutex.Unlock()

	items := make([]T, 0, len(coll.Items))
	for _, a := range coll.Items {
		items = append(items, a)
	}
	return items
}

func (coll *Collection[T]) Create(data []byte) (T, error) {
	var item T
	coll.Mutex.Lock()
	defer coll.Mutex.Unlock()

	err := json.Unmarshal(data, &item)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return item, err
	}

	// if type T has a field ListenedAt or CreatedAt, set it to the current time
	// if t, ok := item.(interface{ SetListenedAt(string) }); ok {
	// 	t.SetListenedAt(time.Now().Format(time.RFC3339))
	// }

	item.SetCreatedAt(time.Now())

	coll.CurrentID++
	item.SetID(coll.CurrentID)
	coll.Items[item.GetID()] = item

	log.Printf("Item created: %v", item.GetCreatedAt().Unix())

	return item, nil
}

func (coll *Collection[T]) Update(id int, data []byte) (T, error) {
	coll.Mutex.Lock()
	defer coll.Mutex.Unlock()

	item, ok := coll.Items[id]
	if !ok {
		log.Println("Item not found")
		return item, errors.New("not found")
	}

	err := json.Unmarshal(data, &item)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return item, err
	}

	item.SetID(id)
	coll.Items[id] = item

	return item, nil
}
