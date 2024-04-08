package main

import (
	"fmt"
	"ztm_scanner/mappers"
	"ztm_scanner/providers/web"
	"ztm_scanner/ztm"
)

func main() {
	jm := mappers.JsonMapper[[]ztm.Stop]{}
	sp := web.StopsProvider{Url: "https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json"}
	bp := web.ResponseBodyProvider{Url: "https://ckan2.multimediagdansk.pl/departures?stopId="}
	stops := make([]ztm.Stop, 0)
	_ = jm.MapValue(&sp, &stops)
	newStops := ztm.FilterStopsByName("bajana", stops)
	fmt.Println(newStops)
	sm := mappers.JsonMapper[ztm.Schedule]{}
	s := ztm.Schedule{}
	bp.Url = bp.Url + fmt.Sprint(newStops[2].StopId)
	_ = sm.MapValue(&bp, &s)
	fmt.Println(s)

}
