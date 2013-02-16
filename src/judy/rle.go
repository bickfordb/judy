package judy

import "io"
import "encoding/binary"

// Run Length Encoding
// Stream of varints like Off^n On^n Off^n

func WriteRLE(j *Judy1, writer io.Writer) {
  var ones bool = false
  var i, last uint64
  var e bool = true

  for e {
    if ones {
      i, e = j.NextEmpty(i)
    } else {
      i, e = j.Next(i)
    }
    if !e {
      break
    }
    writeUvarint(writer, i - last)
    ones = !ones
    last = i
  }
}

func writeUvarint(w io.Writer, x uint64) (err error) {
  var buf [binary.MaxVarintLen64]byte
  n := binary.PutUvarint(buf[:], x)
  _, err = w.Write(buf[0:n])
  return
}

func ReadRLE(reader io.ByteReader) (j *Judy1, err error) {
  j = NewJudy1()
  var idx uint64
  ones := false
  for {
    var width uint64
    width, err = binary.ReadUvarint(reader)
    if err == io.EOF {
      err = nil
      return
    } else if err != nil {
      j.Free()
      j = nil
      return
    }
    if !ones {
      idx += width
    } else {
      var i uint64
      for i = 0; i < width; i = i + 1 {
        j.Set(idx)
        idx = idx + 1
      }
    }
    ones = !ones
  }
  return
}
