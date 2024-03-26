/*
 * Copyright 2024 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package registry

import "fmt"

const (
	ErrItemNotFoundMessage = "could not find"
)

func NewItemNotFoundError(itemType string, name string) ItemNotFoundError {
	return ItemNotFoundError{
		itemType: itemType,
		name:     name,
		message:  ErrItemNotFoundMessage,
	}
}

type ItemNotFoundError struct {
	itemType string
	name     string
	message  string
}

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("%s %s %s", e.message, e.itemType, e.name)
}
