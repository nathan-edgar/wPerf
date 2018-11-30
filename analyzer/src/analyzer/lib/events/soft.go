package events

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

var (
	sm SoftMap
)

func init() {
	sm = make(SoftMap)
}

type Soft struct {
	CPU   int    `json:"cpu"`
	STime uint64 `json:"stime"`
	ETime uint64 `json:"etime"`
}

type SoftMap map[int]*treemap.Map

func (s *SoftMap) Init(sl []Soft) {
	for _, v := range sl {
		_, ok := (*s)[v.CPU]
		if !ok {
			(*s)[v.CPU] = treemap.NewWith(utils.UInt64Comparator)
		}
		(*s)[v.CPU].Put(v.ETime, v)
	}
}

func LoadSoft(file string) []Soft {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("LoadSoft: ", err)
	}

	es := make([]Soft, 0)
	ls := strings.Split(string(bs), "\n")
	for _, l := range ls {
		var e Soft
		err := json.Unmarshal([]byte(l), &e)
		if err != nil {
			log.Fatalln("LoadSoft: ", err)
		}
		es = append(es, e)
	}

	return es
}

func InitSoftContainer(sl []Soft) {
	sm.Init(sl)
}

func GetCeilingSoft(cpu int, time uint64) *Soft {
	_, v := sm[cpu].Ceiling(time)
	return v.(*Soft)
}
