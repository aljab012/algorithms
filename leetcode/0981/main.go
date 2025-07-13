package main

type TimeValue struct {
	v string
	t int
}
type TimeMap struct {
	store map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], TimeValue{v: value, t: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.store[key]
	if !ok {
		return ""
	}
	l, r := 0, len(values)-1

	result := ""
	for l <= r {
		mid := l + (r-l)/2

		if values[mid].t <= timestamp {
			result = values[mid].v
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
