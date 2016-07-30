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

package controllers

import (
	"container/list"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"


	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"samples/WebIM/models"
	"samples/analyser"
	"os"
	"strings"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"path"
	"encoding/base64"
)

type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{0, ep, user, time.Now().Unix(), msg}
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan string, 10)
	// Send events here to publish them.
	publish = make(chan models.Event, 10)
	// Long polling waiting list.
	waitingList = list.New()
	subscribers = list.New()
)

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
				if (event.User == "seller") {
					inputFile := "/tmp/output.txt"
					writeToFile(event, inputFile)
					out := analyser.AnalyseDependencies(inputFile)
					updateAttributes(out)
				}
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "") // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

func updateAttributes(stdout string) {
	li := list.New()
	deps := strings.Split(stdout, "\n\n")[2]
	fmt.Printf("%v\n", reflect.TypeOf(deps))
	for _, dep := range strings.Split(deps, "\n") {
		subs, _ := regexp.MatchString("^([A-Za-z_\\-]+)\\(([A-Za-z\\-]+)\\-([0-9]+), ([A-Za-z_\\-]+)\\-([0-9]+)\\)$", dep)
		r, _ := regexp.Compile("^([A-Za-z_\\-]+)\\(([A-Za-z\\-]+)\\-([0-9]+), ([A-Za-z_\\-]+)\\-([0-9]+)\\)$")
		matches :=  r.FindStringSubmatch(dep)
		parsedDep := Dependency {Type: matches[1], Members:[2]string{strings.ToLower(matches[2]), strings.ToLower(matches[4])}}
		fmt.Printf("%v\n", subs)
		fmt.Printf("%v\n", parsedDep)
		li.PushBack(parsedDep)
		// ^([A-Za-z_\-]+)\(([A-Za-z\-]+)\-([0-9]+), ([A-Za-z_\-]+)\-([0-9]+)\)$
	}

	for _, attr := range models.SingleAttributes {
		fmt.Printf("%v\n", attr)
		presence := isAttributePresent(li, attr)
		fmt.Printf("Presence of %v: %v\n", attr, presence)
	}

}

func isAttributePresent(li *list.List, attr string) bool {
	for e := li.Front(); e != nil; e = e.Next() {
		dep, ok := (e.Value).(Dependency)
		if ok {
			if (dep.Type == "nsubj" && (dep.Members[0] == attr || dep.Members[1] == attr)) {
				return true
			}

		}
		// do something with e.Value
	}
	return false
}

type Dependency struct {
	Type string
	Members [2]string
}

func writeToFile(event models.Event, inputFile string) {
	fo, err := os.Create(inputFile)
	if err != nil {
			panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
			if err := fo.Close(); err != nil {
					panic(err)
			}
	}()
	if _, err := fo.WriteString(lemmatize(event.Content) + "\n"); err != nil {
			panic(err)
	}
}

func lemmatize(str string) string{
	_, filename, _, _ := runtime.Caller(1)
	f := path.Join(path.Dir(filename), "../helpers/lemmatizer.rb")
	// data := []byte(strings.ToLower(str))
	str = base64.StdEncoding.EncodeToString([]byte("there are many shops"))
	fmt.Printf("%v\n", str)
	return analyser.PrintAndExec(fmt.Sprintf("ruby %v %v base64", f, str))
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
