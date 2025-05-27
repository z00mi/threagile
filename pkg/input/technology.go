package input

import "fmt"

type Technology struct {
	Kind        string `yaml:"kind,omitempty" json:"kind,omitempty"`
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	Version     string `yaml:"version,omitempty" json:"version,omitempty"`
}

func (what *Technology) Merge(other Technology) error {
	var mergeError error
	what.Kind, mergeError = new(Strings).MergeSingleton(what.Kind, other.Kind)
	if mergeError != nil {
		return fmt.Errorf("failed to merge name: %w", mergeError)
	}

	what.Name, mergeError = new(Strings).MergeSingleton(what.Name, other.Name)
	if mergeError != nil {
		return fmt.Errorf("failed to merge name: %w", mergeError)
	}

	what.Description, mergeError = new(Strings).MergeSingleton(what.Description, other.Description)
	if mergeError != nil {
		return fmt.Errorf("failed to merge description: %w", mergeError)
	}

	what.Version, mergeError = new(Strings).MergeSingleton(what.Version, other.Version)
	if mergeError != nil {
		return fmt.Errorf("failed to merge version: %w", mergeError)
	}

	return nil
}

func (what *Technology) MergeMap(first map[string]Technology, second map[string]Technology) (map[string]Technology, error) {
	for mapKey, mapValue := range second {
		mapItem, ok := first[mapKey]
		if ok {
			mergeError := mapItem.Merge(mapValue)
			if mergeError != nil {
				return first, fmt.Errorf("failed to merge technology %q: %w", mapKey, mergeError)
			}

			first[mapKey] = mapItem
		} else {
			first[mapKey] = mapValue
		}
	}

	return first, nil
}
