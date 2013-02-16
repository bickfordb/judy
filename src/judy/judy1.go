package judy

// #cgo LDFLAGS: -lJudy
// #include <Judy.h>
import "C"
import "unsafe"
import "fmt"

type Judy1 struct {
  val unsafe.Pointer
}

func NewJudy1() *Judy1 {
  return &Judy1{}
}

func (j *Judy1) Set(bit uint64) {
  C.Judy1Set(C.PPvoid_t(&j.val), C.Word_t(bit), nil)
}

func (j *Judy1) Unset(bit uint64) {
  C.Judy1Unset(C.PPvoid_t(&j.val), C.Word_t(bit), nil)
}

func (j *Judy1) Count(index1 uint64, index2 uint64) uint64 {
  return uint64(C.Judy1Count(C.Pcvoid_t(j.val), C.Word_t(index1), C.Word_t(index2), nil))
}

func (j *Judy1) ByCount(count uint64) (index uint64, found bool) {
  var ret C.Word_t = 0
  st := C.Judy1ByCount(C.Pcvoid_t(j.val), C.Word_t(count), &ret, nil)
  if st == 1 {
    found = true
    index = uint64(ret)
  }
  return
}

func (j *Judy1) Free() {
  C.Judy1FreeArray(&j.val, nil)
}

func (j *Judy1) Test(index uint64) (exists bool) {
  st := C.Judy1Test(C.Pcvoid_t(j.val), C.Word_t(index), nil)
  if st != 0 {
    exists = true
  }
  return
}

func (j *Judy1) First(index uint64) (firstIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1First(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    firstIndex = uint64(index0)
    found = true
  }
  return
}

func (j *Judy1) FirstEmpty(index uint64) (firstIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1FirstEmpty(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    firstIndex = uint64(index0)
    found = true
  }
  return
}
func (j *Judy1) Next(index uint64) (nextIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1Next(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    nextIndex = uint64(index0)
    found = true
  }
  return
}

func (j *Judy1) NextEmpty(index uint64) (nextIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1NextEmpty(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    nextIndex = uint64(index0)
    found = true
  }
  return
}

func (j *Judy1) Last(index uint64) (lastIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1Prev(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    lastIndex = uint64(index0)
    found = true
  }
  return
}

func (j *Judy1) Prev(index uint64) (prevIndex uint64, found bool) {
  index0 := C.Word_t(index)
  st := C.Judy1Prev(C.Pcvoid_t(j.val), &index0, nil)
  if st != 0 {
    prevIndex = uint64(index0)
    found = true
  }
  return
}

func (j *Judy1) Copy() (aCopy *Judy1) {
  aCopy = NewJudy1()
  for i, exists := j.First(0); exists; i, exists = j.Next(i) {
    aCopy.Set(i)
  }
  return
}

func (j *Judy1) Equals(other *Judy1) bool {
  for i, e := j.First(0); e; i, e = j.Next(i) {
    if !other.Test(i) { return false }
  }
  for i, e := other.First(0); e; i, e = other.Next(i) {
    if !j.Test(i) { return false }
  }
  return true
}

func (j *Judy1) Repr() string {
  s := "Judy1("
  for i, e := j.First(0); e; i, e = j.Next(i) {
    if s != "Judy1(" {
      s = s + ", "
    }
    s = s + fmt.Sprintf("%v", i)
  }
  s += ")"
  return s
}


