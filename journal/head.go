package journal

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Get Head size
// Put Head item
// UpdateHead takes a key and adds it to the head key
// TODO: Review and optimize
func (j *Journal) UpdateHead(key string) error {
	// We need to sort and pop oldest entry
	keys := j.GetJournalHeadKeys()
	size := len(keys)
	if size < 5 {
		// append
		keys = append(keys, key)
		// Prep for adding to db
		sort.Strings(keys)
		keyString := strings.Join(keys, " ")
		err := j.DB.Put([]byte("head"), []byte(keyString))
		if err != nil {
			return err
		}
	} else {
		// Replace
		sort.Strings(keys)
		keys[0] = key
		sort.Strings(keys)
		// Prep for adding to db
		keyString := strings.Join(keys, " ")
		err := j.DB.Put([]byte("head"), []byte(keyString))
		if err != nil {
			return err
		}
	}
	j.HeadKeys = keys

	return nil
}

func (j *Journal) Head() error {
	payload := make(map[int64]Message)
	for _, key := range j.HeadKeys {
		// Getting string value
		val, err := j.DB.Get([]byte(key))
		if err != nil {
			return err
		}
		var msg Message
		err = json.Unmarshal(val, &msg)
		if err != nil {
			return err
		}
		k, _ := strconv.Atoi(key)
		payload[int64(k)] = msg
	}
	j.printJournalValues(payload)
	return nil
}

// GetJournalHeadKeys pull a list of head keys and serializes them to a list of strings
func (j *Journal) GetJournalHeadKeys() []string {
	val, err := j.DB.Get([]byte("head"))
	if err != nil {
		log.Fatalln("error setting head", err)
	}
	// new db/no head keys
	if val == nil {
		return []string{}
	}

	return strings.Split(string(val), " ")
}

//printJournalValues displays journal output
func (j *Journal) printJournalValues(payload map[int64]Message) {
	for key, msg := range payload {
		ts := KeyToTime(key, time.Kitchen)
		fmt.Printf("[ %s ]: %s\n", ts, msg.Content)
	}
}

func KeyToTime(key int64, timeFormat string) string {
	t := time.Unix(key, 0).UTC().Format(timeFormat)
	return t
}
