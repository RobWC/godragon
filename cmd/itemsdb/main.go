package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/robwc/godragon"
)

func updateDatabase(db *bolt.DB, version string) error {

	items, err := godragon.StaticItems(version)
	if err != nil {
		return err
	}
	var buff bytes.Buffer

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("items"))
		return err
	})

	for i := range items {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("items"))

			enc := gob.NewEncoder(&buff)
			enc.Encode(items[i])
			err := b.Put([]byte(items[i].Name), buff.Bytes())
			buff.Reset()
			return err
		})
	}
	return nil
}

func readItems(db *bolt.DB) map[string]godragon.Item {

	itemm := make(map[string]godragon.Item)

	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("items")).Cursor()

		var buff bytes.Buffer

		for k, v := c.First(); k != nil; k, v = c.Next() {

			var item godragon.Item
			dec := gob.NewDecoder(&buff)
			buff.Write(v)
			err := dec.Decode(&item)
			if err != nil {
				log.Println(err)
				return err
			}
			itemm[string(k)] = item
			buff.Reset()
		}
		return nil
	})

	return itemm

}

func main() {
	db, err := bolt.Open("items.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	updateDatabase(db, "6.4.1")
	items := readItems(db)

	for i := range items {
		if items[i].Stats.FlatPhysicalDamageMod != 0 || items[i].Stats.PercentPhysicalDamageMod != 0 {
			fmt.Printf("%-40s %-4d %-8f %-8f\n", items[i].Name, items[i].Cost.Total, items[i].Stats.FlatPhysicalDamageMod, items[i].Stats.PercentPhysicalDamageMod)
		}
	}
}
