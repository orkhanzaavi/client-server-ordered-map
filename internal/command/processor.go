package command

import (
	"errors"
	"fmt"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type Storage interface {
	Set(key, value string)
	Get(key string) *StoredItem
	GetAll() []*StoredItem
	Delete(key string)
}

type Processor struct {
	storage Storage
}

func NewProcessor(storage Storage) *Processor {
	return &Processor{
		storage: storage,
	}
}

func NewProcessorWithDefaultStorage() *Processor {
	return &Processor{
		storage: newDefaultStorage(),
	}
}

func (p *Processor) Process(cmd Command) error {
	switch cmd.Name {
	case "test":
		return p.test(cmd)
	case "getItem":
		return p.getItem(cmd)
	case "addItem":
		return p.addItem(cmd)
	case "deleteItem":
		return p.deleteItem(cmd)
	case "getAllItems":
		return p.getAllItems(cmd)
	default:
		return ErrInvalidCommandName
	}
}

func (p *Processor) test(cmd Command) error {
	// Do something
	return nil
}

func (p *Processor) getItem(cmd Command) error {
	item := p.storage.Get(cmd.Arguments[0])
	if item == nil {
		return ErrItemNotFound
	}
	fmt.Printf("%s:%s\n", item.Key, item.Value)
	return nil
}

func (p *Processor) addItem(cmd Command) error {
	p.storage.Set(cmd.Arguments[0], cmd.Arguments[1])
	return nil
}

func (p *Processor) deleteItem(cmd Command) error {
	p.storage.Delete(cmd.Arguments[0])
	return nil
}

func (p *Processor) getAllItems(cmd Command) error {
	items := p.storage.GetAll()
	fmt.Println("All items:")
	for _, item := range items {
		fmt.Printf("%s:%s\n", item.Key, item.Value)
	}

	return nil
}
