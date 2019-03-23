package main

import (
	"encoding/json"
	"fmt"
)

// the following json is provided by @valorad/inventory
var invItems = []byte(`[
  {
    "id": "xmarkthespot",
    "holder": "actor-dragonborn",
    "item": "item-stimpack",
    "refDetails": [
      {
        "refID": "wearewaitingfortes6",
        "owner": "actor-dragonborn",
        "num": 1,
        "item": "item-stimpack"
      },
      {
        "refID": "fallout76pugai",
        "owner": "actor-dragonborn",
        "num": 2,
        "item": "item-stimpack"
      },
      {
        "refID": "sc2terrible",
        "owner": "actor-dragonborn",
        "num": 3,
        "item": "item-stimpack"
      }
    ],
    "base": {
      "dbname": "item-stimpack",
      "name": "Stimpack",
      "description": "",
      "value": 152,
      "weight": 0.1,
      "category": "comsumables",
      "detail": {
        "type": "type-potion",
        "typeName": "Potions",
        "effectsI18n": [
          "Heal"
        ]
      }
    }
  },
  {
    "id": "whosyourdaddy",
    "holder": "actor-dragonborn",
    "item": "item-iron_helmet",
    "refDetails": [
      {
        "refID": "gpangDontknow3",
        "owner": "actor-dragonborn",
        "num": 4,
        "item": "item-iron_helmet"
      }
    ],
    "base": {
      "dbname": "item-iron_helmet",
      "name": "Iron Helmet",
      "description": "",
      "value": 75,
      "weight": 2,
      "category": "gears",
      "detail": {
        "rating": 10,
        "type": "type-heavyarmor",
        "typeName": "Heavy Armor",
        "equipI18n": [
          {
            "name": "head",
            "equip": "equip-head"
          }
        ],
        "effectsI18n": []
      }
    }
  },
  {
    "id": "robinhood",
    "holder": "actor-dragonborn",
    "item": "item-darkbro_tenant",
    "refDetails": [
      {
        "refID": "jadoretesassiettes",
        "owner": "actor-dragonborn",
        "num": 5,
        "item": "item-iron_helmet"
      }
    ],
    "base": {
      "dbname": "item-darkbro_tenant",
      "name": "Dark brother tenant",
      "description": "The divine tenant of Dark brothers",
      "value": 5,
      "weight": 1,
      "category": "books",
      "detail": {
        "content": "content-darkbro_tenant",
        "contentDetail": "Dark-bro rules: 1. No black sacrament on kittens. 2. No D.."
      }
    }
  }
]`)

// InvItemSlice {invItem[]}
type InvItemSlice struct {
	Items []InvItem
}

// InvItem {invItem}
type InvItem struct {
	Id   string
	Item string
}

func unMarshalToStruct() InvItemSlice {
	var invItemslice InvItemSlice
	err := json.Unmarshal(invItems, &invItemslice.Items)
	fmt.Println(err)
	return invItemslice
}

func main() {
	fmt.Println(unMarshalToStruct())
}
