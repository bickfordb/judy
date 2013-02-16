package judy

import _ "fmt"
import "testing"
import "bytes"

func TestJudy(t *testing.T) {
  judy := NewJudy1()
  defer judy.Free()
  judy.Set(5)
  if !judy.Test(5) {
    t.Errorf("expecting 5 to be set")
    t.FailNow()
  }
  judy.Unset(5)
  if judy.Test(5) {
    t.Errorf("unexpected 5 is set")
    t.FailNow()
  }

  judy.Set(10)
  judy.Set(20)

  if i, e := judy.First(0); !e || i != 10 {
    t.Fatalf("Expecting first index to be 10, %s, %s", i, e)
  }
  if i, e := judy.Next(10); !e || i != 20 {
    t.Fatalf("Expecting next index to be 20, %s, %s", i, e)
  }
  if i, e := judy.Last(21); !e || i != 20 {
    t.Fatalf("Expecting last index to be 20, %s, %s", i, e)
  }
  if i, e := judy.Prev(20); !e || i != 10 {
    t.Fatalf("Expecting prev index to be 10, %s, %s", i, e)
  }
}

func TestCopy(t *testing.T) {
  j := NewJudy1()
  defer j.Free()
  j.Set(5)
  k := j.Copy()
  defer k.Free()
  if !k.Equals(j) {
    t.Fatal("expecting k==j", k.Repr(), j.Repr())
  }
}

func TestRLE(t *testing.T) {
  j := NewJudy1()
  defer j.Free()
  j.Set(5)
  j.Set(6)
  j.Set(10)
  j.Set(11)
  j.Set(13)
  dstBuf := new(bytes.Buffer)
  WriteRLE(j, dstBuf)
  srcBuf := bytes.NewBuffer(dstBuf.Bytes())
  j2, err := ReadRLE(srcBuf)
  if err != nil {
    t.Fatal(err)
  }
  if !j2.Equals(j) {
    t.Fatal("expecting j == j2", j.Repr(), j2.Repr())
  }
}

