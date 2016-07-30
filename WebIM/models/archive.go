// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"container/list"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Id				uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Type      EventType // JOIN, LEAVE, MESSAGE
	User      string
	Timestamp int64// Unix timestamp (secs)
	Content   string
}

type Attribute struct {
	Id				uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Known 		bool
	Presence  bool // JOIN, LEAVE, MESSAGE
	Value     string
	Created 	int64// Unix timestamp (secs)
	Modified 	int64// Unix timestamp (secs)
}


	var SingleAttributes []string = []string {"hall", "a/c", "fridge", "refrigerator", "parking", "generator", "invertor", "cupboards", "maintenance", "tv", "beds", "lift", "floor" }
	var CompoundAttributes []string = []string {"modular kitchen"}

func init() {
	db, _ := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Attribute{})
}
const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
	db, _ := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
	db.Create(&event)
}

// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int64(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
