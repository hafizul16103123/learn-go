package main
import(
	"fmt"
)

/*
Shallow Copy:
	Shallow Copy মানে নতুন variable তৈরি হয়, কিন্তু ভিতরের reference (pointer, slice, map ইত্যাদি) একই memory-কে point করে।
	দুইটি slice-এর header আলাদা, কিন্তু underlying array একই।
	=Slice assignment Shallow Copy।
	=Map assignment Shallow Copy।



Deep Copy:
	Deep Copy মানে সম্পূর্ণ নতুন memory তৈরি হয়।
	এখন একটাকে পরিবর্তন করলে অন্যটি পরিবর্তিত হবে না।
	=Array assignment Deep Copy করে।

| Type                                | Default Copy | Shared Memory?                |
| ----------------------------------- | ------------ | ----------------------------- |
| Array                               | Deep Copy    | ❌ No                          |
| Slice                               | Shallow Copy | ✅ Yes (underlying array)      |
| Map                                 | Shallow Copy | ✅ Yes                         |
| Pointer                             | Shallow Copy | ✅ Yes                         |
| Channel                             | Shallow Copy | ✅ Yes                         |
| Struct (only value fields)          | Deep Copy    | ❌ No                          |
| Struct (contains slice/map/pointer) | Mixed        | ✅ Reference fields are shared |


Interview Shortcut (30 seconds):
✅ Deep Copy: Basic types, Arrays, Structs (only value fields)
❌ Shallow Copy: Slices, Maps, Pointers, Channels
⚠️ Mixed: Structs containing slices/maps/pointers
✅ make()+copy() gives a true deep copy only when the elements are value types. If the elements themselves contain references, you'll need to deep-copy those recursively.


"Are slices deep copied in Go?"

না। Slice assignment শুধুমাত্র slice header (pointer, length, capacity) copy করে। Underlying array share হয়। Deep copy করতে make() এবং copy() ব্যবহার করতে হয়।

*/

func main(){
	s:="hello"
	s2:=s
	s2="World"
	fmt.Println(&s)
	fmt.Println(&s2)

	a:=1
	a2:=a
	a2=2
	fmt.Println(&a)
	fmt.Println(&a2)

	sl1:=[]int{1,2}
	sl2:=sl1
	sl2[0]=9
	fmt.Println(&sl1[0])
	fmt.Println(&sl2[0])
}